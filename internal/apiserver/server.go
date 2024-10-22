package apiserver

import (
	"errors"
	"fmt"
	"net/http"
	"obstore/internal/model"
	"obstore/internal/store"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not autenticated")
)

var (
	hmacSampleSecret = "8a046a6b436496d9c7af3e196a73ee9948677eb30b251706667ad59d6261bd78d2f6f501a6dea0118cfb3b0dcd62d6c9eb88142e2c24c2c686133a935cd65651"
)

type server struct {
	router *gin.Engine
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: gin.Default(),
		store:  store,
	}

	s.router.Use(gin.Logger())
	/*
		s.router.Use(cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST"},
				AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
				ExposeHeaders:    []string{"Content-Length", "application/json"},
			AllowCredentials: true,
				AllowOriginFunc: func(origin string) bool {
					return origin == "http://localhost:8090/view"
				},
			MaxAge: 24 * time.Hour,
		}))
	*/

	// Настройка CORS
	s.router.Use(CORSMiddleware())

	s.configureRouter()

	return s
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (s *server) configureRouter() {
	s.router.Static("/css", "./frontend/css/")
	s.router.StaticFS("/scripts", http.Dir("./frontend/scripts"))
	s.router.StaticFile("/favicon.png", "./fronend/warehouse.png")
	s.router.LoadHTMLGlob("frontend/**/*")

	// Маршруты для API
	apiGroup := s.router.Group("/api/v1")
	{
		userGroup := apiGroup.Group("/user", s.AuthMW)
		{
			userGroup.POST("/signout", s.SignOutUserById)
			userGroup.POST("/update", s.UpdateUser)
			userGroup.POST("/update/pass", s.UpdatePassword)
		}

		usersGroup := apiGroup.Group("/users")
		{
			usersGroup.POST("", s.CreateUser)
			usersGroup.POST("/signin", s.SignIn)
			usersGroup.GET("", s.AuthMW, s.GetUserAll)
			usersGroup.POST("/password/restore", s.SetUserTemporaryPassword)
		}

		teamGroup := apiGroup.Group("/team", s.AuthMW)
		{
			teamGroup.GET("", s.GetTeamById)
			teamGroup.POST("/delete", s.DeteleTeamById)
			teamGroup.POST("/update", s.UpdateTeam)
			teamGroup.POST("/teamcomposition", s.GetTeamComposition)

		}

		teamsGroup := apiGroup.Group("/teams", s.AuthMW)
		{
			teamsGroup.POST("", s.CreateTeam)
			teamsGroup.GET("", s.GetTeamAll)
		}

		productGroup := apiGroup.Group("/product", s.AuthMW)
		{
			productGroup.GET("/find/article", s.GetProductByArticle)
			productGroup.GET("/find/strikecode", s.GetProductByStrikeCode)
			productGroup.GET("/find/name", s.GetProductByName)
			productGroup.POST("/update/strikecode", s.UpdateProductStrikeCodeById)
		}

		productsGroup := apiGroup.Group("/products", s.AuthMW)
		{
			productsGroup.POST("", s.CreateProduct)
			productsGroup.GET("", s.GetProductsAll)
		}

		orderGroup := apiGroup.Group("/order", s.AuthMW)
		{
			orderGroup.GET("/find/id", s.GetOrderById)
			orderGroup.GET("/find/uid", s.GetOrderByUID)
			orderGroup.GET("/find/num", s.GetOrderByFolioNum)
		}

		ordersGroup := apiGroup.Group("/orders", s.AuthMW)
		{
			ordersGroup.POST("", s.CreateOrder)
			ordersGroup.GET("", s.GetOrdersAll)
			ordersGroup.GET("/range", s.GetOrdersByDateRange)
			ordersGroup.GET("/driver", s.GetOrdersByDriver)
			ordersGroup.GET("/agent", s.GetOrdersByAgent)
		}

		teamCompositionGroup := apiGroup.Group("/teamcomposition", s.AuthMW)
		{
			teamCompositionGroup.GET("/", s.GetTeamCompositionById)
			teamCompositionGroup.POST("/update", s.UpdateTeamComposition)
			teamCompositionGroup.GET("/team", s.GetTeamCompositionByTeamId)
			teamCompositionGroup.GET("/user", s.GetTeamCompositionByUserId)
		}

		teamCompositionsGroup := apiGroup.Group("/teamcompositions", s.AuthMW)
		{
			teamCompositionsGroup.POST("", s.CreateTeamComposition)
			teamCompositionsGroup.GET("", s.GetTeamCompositionAll)

		}

		assemblyOrderGroup := apiGroup.Group("/assemblyorder", s.AuthMW)
		{
			assemblyOrderGroup.GET("/find/id", s.GetAssemblyOrderByID)
		}

		assemblyOrdersGroup := apiGroup.Group("/assemblyorders", s.AuthMW)
		{
			assemblyOrdersGroup.POST("", s.CreateAssemblyOrder)
		}

		warehousesGroup := apiGroup.Group("/warehouses", s.AuthMW)
		{
			warehousesGroup.GET("", s.GetWarehouseAll)
			warehousesGroup.POST("", s.CreateWarehouse)

		}

		warehouseGroup := apiGroup.Group("/warehouse", s.AuthMW)
		{
			warehouseGroup.GET("/find/id", s.GetWarehouseById)
			warehouseGroup.POST("/update", s.UpdateWarehouse)
			warehouseGroup.POST("/delete", s.DeleteWarehouseById)
		}

		rolesGroup := apiGroup.Group("/roles")
		{
			rolesGroup.POST("", s.CreateRole)
			rolesGroup.GET("", s.GetAllRoles)
		}
		roleGroup := apiGroup.Group("/role")
		{
			roleGroup.POST("/update", s.UpdateRole)
			roleGroup.POST("/delete", s.DeleteRole)
			roleGroup.GET("/find/id", s.GetRoleById)
			roleGroup.GET("/find/idrole", s.GetRoleUser)
		}

		userRolesGroup := apiGroup.Group("/userRoles")
		{
			userRolesGroup.POST("", s.CreateUserRole)
			userRolesGroup.GET("", s.GetAllUserRole)
			userRolesGroup.GET("/find/userId", s.GetUserRoleByUserId)
			userRolesGroup.GET("/find/roleId", s.GetUserRoleByRoleId)
		}
		userRoleGroup := apiGroup.Group("/userRole")
		{
			//userRoleGroup.POST("/delete", s.DeleteUserRole)
			userRoleGroup.POST("/update", s.UpdateUserRole)
			userRoleGroup.GET("/find/id", s.GetUserRoleById)

		}

	}

	/*
		// Маршруты для сайта
		viewGroup := s.router.Group("/view")
		{
			viewGroup.GET("/login", s.LoginHTML)

			userGroup := viewGroup.Group("/user")
			{
				userGroup.POST("/signout", s.SignOutUserById)
				userGroup.POST("/update", s.UpdateUser)
				userGroup.POST("/update/pass", s.UpdatePassword)
			}

			usersGroup := userGroup.Group("/users")
			{
				usersGroup.POST("", s.CreateUser)
				usersGroup.POST("/signin", s.SignIn)
			}
		}

	*/
}

// User..
func (s *server) CreateUser(ctx *gin.Context) {
	var user model.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err = s.store.User().Create(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (s *server) SignIn(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(user.Email, user.Pass)

	user, err = s.store.User().SignInUser(user.Email, user.Pass)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errIncorrectEmailOrPassword.Error()})
		return
	}

	tokenString, err := createAndSignJWT(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JWT creation failed. Error: " + err.Error()})
		return
	}

	err = s.store.User().UpdateToken(user.ID, tokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed update JWT user. Error: " + err.Error()})
		return
	}

	user.Token = tokenString
	setCookie(ctx, tokenString)
	ctx.JSON(http.StatusOK, user)
}

func (s *server) SignOutUserById(ctx *gin.Context) {
	type request struct {
		Id int `json:"id"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = s.store.User().SignOutUserById(req.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("Auth", "deleted", 0, "", "", false, false)

	ctx.JSON(http.StatusAccepted, gin.H{"message": "success"})
}

func (s *server) UpdateUser(ctx *gin.Context) {
	var u model.User
	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err = s.store.User().Update(u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, u)
}

func (s *server) UpdatePassword(ctx *gin.Context) {
	type request struct {
		Id   int    `json:"id" validate:"required"`
		Pass string `json:"pass" validate:"required"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = s.store.User().ChangePassword(req.Id, req.Pass)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

func (s *server) GetUserAll(ctx *gin.Context) {
	users, err := s.store.User().GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (s *server) SetUserTemporaryPassword(ctx *gin.Context) {
	type request struct {
		Email string `json:"email"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pass, err := s.store.User().SetTemporaryPassword(req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"password": pass})
}

func createAndSignJWT(user *model.User) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"ttl":    time.Now().Add(time.Hour * 24 * 100).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(hmacSampleSecret))
}

func setCookie(ctx *gin.Context, token string) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Auth", token, 3600*24*100, "", "", false, true)
}

func (s *server) AuthMW(ctx *gin.Context) {
	// Получение токена из куки
	tokenStr, err := ctx.Cookie("Auth")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No auth token"})
		ctx.Abort()
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(hmacSampleSecret), nil
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse JWT"})
		ctx.Abort()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "JWT Claims failed"})
		ctx.Abort()
	}

	if claims["ttl"].(float64) < float64(time.Now().Unix()) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "JWT token expired"})
		ctx.Abort()
	}

	user, err := s.store.User().UserFromID(claims["userID"].(float64))

	if user.ID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Could not find the user!"})
		ctx.Abort()
	}

	ctx.Set("user", user)

	ctx.Next()
}

// Team...

func (s *server) CreateTeam(ctx *gin.Context) {
	var teams []model.Team

	err := ctx.ShouldBindJSON(&teams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var createdTeams []model.Team
	for _, team := range teams {
		team, err = s.store.Team().Create(team)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		} else {
			createdTeams = append(createdTeams, team)
		}
	}

	ctx.JSON(http.StatusCreated, createdTeams)
}

func (s *server) GetTeamAll(ctx *gin.Context) {
	teams, err := s.store.Team().GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, teams)
}

func (s *server) GetTeamById(ctx *gin.Context) {
	type request struct {
		Id int `json:"id" validate:"required"`
	}

	var reqs []request
	err := ctx.ShouldBindJSON(&reqs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var findedTeams []model.Team

	for _, req := range reqs {
		team, err := s.store.Team().GetById(req.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		} else {
			findedTeams = append(findedTeams, team)
		}
	}

	ctx.JSON(http.StatusOK, findedTeams)
}

func (s *server) DeteleTeamById(ctx *gin.Context) {
	type request struct {
		Id int `json:"id" validate:"required"`
	}

	var reqs []request
	err := ctx.ShouldBindJSON(&reqs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, req := range reqs {
		err := s.store.Team().DeleteById(req.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Запрос на удаление сборочных команд успешно выполнен"})
}

func (s *server) UpdateTeam(ctx *gin.Context) {
	var team model.Team
	err := ctx.ShouldBindJSON(&team)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err = s.store.Team().Update(team)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, team)
}

func (s *server) GetTeamComposition(ctx *gin.Context) {
	type request struct {
		ID uint `json:"id"`
	}

	var reqs []request
	err := ctx.ShouldBindJSON(&reqs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var findedTC []model.Team
	for _, req := range reqs {
		fmt.Println("ID - ", req.ID)

		tc, err := s.store.Team().TeamComposition(req.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			continue
		} else {
			findedTC = append(findedTC, tc)
		}
	}

	ctx.JSON(http.StatusOK, findedTC)
}

// Product ...

func (s *server) CreateProduct(ctx *gin.Context) {
	var products []model.Product

	err := ctx.ShouldBindJSON(&products)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var createdProducts []model.Product
	for _, product := range products {
		product, err = s.store.Product().Create(product)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			createdProducts = append(createdProducts, product)
		}
	}

	ctx.JSON(http.StatusCreated, createdProducts)
}

func (s *server) GetProductsAll(ctx *gin.Context) {
	products, err := s.store.Product().GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (s *server) GetProductByArticle(ctx *gin.Context) {
	type request struct {
		Article string `json:"article"`
	}

	var reqs []request
	err := ctx.ShouldBindJSON(&reqs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var findedProduct []model.Product
	for _, req := range reqs {
		product, err := s.store.Product().GetByArticle(req.Article)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			findedProduct = append(findedProduct, product)
		}
	}

	ctx.JSON(http.StatusOK, findedProduct)
}

func (s *server) GetProductByStrikeCode(ctx *gin.Context) {
	var product model.Product

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	findedProduct, err := s.store.Product().GetByStrikeCode(product.StrikeCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, findedProduct)
}

func (s *server) GetProductByName(ctx *gin.Context) {
	type request struct {
		NameArtic string `json:"name_artic"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	findedProduct, err := s.store.Product().GetByName(req.NameArtic)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, findedProduct)
}

func (s *server) UpdateProductStrikeCodeById(ctx *gin.Context) {
	type request struct {
		Id         int    `json:"id" validate:"required"`
		StrikeCode string `json:"strikecode" validate:"required"`
	}

	var reqs []request
	err := ctx.ShouldBindJSON(&reqs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedProduct []model.Product
	for _, req := range reqs {
		product, err := s.store.Product().UpdateStrikeCode(req.Id, req.StrikeCode)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		} else {
			updatedProduct = append(updatedProduct, product)
		}
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}

// Orders ...

func (s *server) CreateOrder(ctx *gin.Context) {
	var reqs []model.Order

	err := ctx.ShouldBindJSON(&reqs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var createdOrders []model.Order
	for _, req := range reqs {
		order, err := s.store.Order().Create(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		} else {
			createdOrders = append(createdOrders, order)
		}
	}

	ctx.JSON(http.StatusCreated, createdOrders)
}

func (s *server) GetOrderById(ctx *gin.Context) {
	type request struct {
		Id int `json:"id"`
	}

	var reqs []request
	err := ctx.ShouldBindJSON(&reqs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var orders []model.Order

	for _, req := range reqs {
		order, err := s.store.Order().GetById(req.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		} else {
			orders = append(orders, order)
		}
	}

	ctx.JSON(http.StatusOK, orders)
}

func (s *server) GetOrderByUID(ctx *gin.Context) {
	type request struct {
		OrderUID int `json:"order_uid" validate: "required"`
	}

	var reqs []request
	err := ctx.ShouldBindJSON(&reqs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var orders []model.Order

	for _, req := range reqs {
		order, err := s.store.Order().GetByOrderUID(req.OrderUID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		} else {
			orders = append(orders, order)
		}
	}

	ctx.JSON(http.StatusOK, orders)
}

func (s *server) GetOrderByFolioNum(ctx *gin.Context) {
	type request struct {
		FolioNum int `json:"folio_num" validate: "required"`
	}

	var reqs []request
	err := ctx.ShouldBindJSON(&reqs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var orders []model.Order

	for _, req := range reqs {
		order, err := s.store.Order().GetByFolioNum(req.FolioNum)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		} else {
			orders = append(orders, order)
		}
	}

	ctx.JSON(http.StatusOK, orders)
}

func (s *server) GetOrdersAll(ctx *gin.Context) {
	orders, err := s.store.Order().GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (s *server) GetOrdersByDriver(ctx *gin.Context) {
	type request struct {
		Driver string `json:"driver"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	findedOrders, err := s.store.Order().GetByDriver(req.Driver)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, findedOrders)
}

func (s *server) GetOrdersByAgent(ctx *gin.Context) {
	type request struct {
		Agent string `json:"agent"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	findedOrders, err := s.store.Order().GetByAgent(req.Agent)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, findedOrders)
}

func (s *server) GetOrdersByDateRange(ctx *gin.Context) {
	type request struct {
		DtStart  string `json:"dt_start"`
		DtFinish string `json:"dt_finish"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	findedOrders, err := s.store.Order().GetByDateRange(req.DtStart, req.DtFinish)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, findedOrders)
}

// Team Compositions

func (s *server) CreateTeamComposition(ctx *gin.Context) {
	var reqs []model.TeamComposition

	err := ctx.ShouldBindJSON(&reqs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var createdTS []model.TeamComposition
	for _, req := range reqs {
		ts, err := s.store.TeamComposition().Create(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		} else {
			createdTS = append(createdTS, ts)
		}
	}

	ctx.JSON(http.StatusCreated, createdTS)
}

func (s *server) GetTeamCompositionById(ctx *gin.Context) {
	type request struct {
		ID int `json:"id"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tsc, err := s.store.TeamComposition().GetByID(uint(req.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tsc)
}

func (s *server) UpdateTeamComposition(ctx *gin.Context) {
	var ts model.TeamComposition
	err := ctx.ShouldBindJSON(&ts)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ts, err = s.store.TeamComposition().Update(ts)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, ts)
}

func (s *server) GetTeamCompositionByTeamId(ctx *gin.Context) {
	type request struct {
		ID int `json:"id"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ts, err := s.store.TeamComposition().GetByTeamId(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, ts)
}

func (s *server) GetTeamCompositionAll(ctx *gin.Context) {
	tcs, err := s.store.TeamComposition().GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, tcs)
}

func (s *server) GetTeamCompositionByUserId(ctx *gin.Context) {
	type request struct {
		UserID int `json:"user_id"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tcs, err := s.store.TeamComposition().GetByUserId(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tcs)
}

// Site ...
func (s *server) LoginHTML(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"title": "ТД Восток"})
}

// AssemblyOrder ...
func (s *server) GetAssemblyOrderByID(ctx *gin.Context) {
	type request struct {
		ID uint `json:"id"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ao, err := s.store.AssemblyOrder().GetByID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, ao)
}

func (s *server) CreateAssemblyOrder(ctx *gin.Context) {
	var ao model.AssemblyOrder
	err := ctx.ShouldBindJSON(&ao)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Удаляем запись по ID
	ao, err = s.store.AssemblyOrder().Create(ao)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, ao)
}

// Warehiuses
func (s *server) GetWarehouseAll(ctx *gin.Context) {
	warehouse, err := s.store.Warehouse().GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, warehouse)
}

func (s *server) GetWarehouseById(ctx *gin.Context) {
	type request struct {
		ID uint `json:"id"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(req.ID)

	wh, err := s.store.Warehouse().GetByID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, wh)
}

func (s *server) CreateWarehouse(ctx *gin.Context) {
	var reqs []model.Warehouse

	err := ctx.ShouldBindJSON(&reqs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var createdWarhouse []model.Warehouse
	for _, req := range reqs {
		warehouse, err := s.store.Warehouse().Create(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		} else {
			createdWarhouse = append(createdWarhouse, warehouse)
		}

	}
	ctx.JSON(http.StatusCreated, createdWarhouse)
}

func (s *server) UpdateWarehouse(ctx *gin.Context) {
	var warehouse model.Warehouse
	err := ctx.ShouldBindJSON(&warehouse)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Обновляем склад
	updatedWarehouse, err := s.store.Warehouse().Update(warehouse)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем обновленный объект
	ctx.JSON(http.StatusOK, updatedWarehouse)
}

func (s *server) DeleteWarehouseById(ctx *gin.Context) {
	type request struct {
		Id uint `json:"id"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Удаляем запись по ID
	err = s.store.Warehouse().DeleteByID(req.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Запись успешно удалена"})
}

//

func (s *server) CreateRole(ctx *gin.Context) {
	var roles []model.Role

	err := ctx.ShouldBindJSON(&roles)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var createdRoles []model.Role
	for _, role := range roles {
		role, err = s.store.Role().Create(role)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		} else {
			createdRoles = append(createdRoles, role)
		}
	}
	ctx.JSON(http.StatusCreated, createdRoles)
}

func (s *server) GetAllRoles(ctx *gin.Context) {
	roles, err := s.store.Role().All()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, roles)
}

func (s *server) GetRoleById(ctx *gin.Context) {
	type request struct {
		Id uint `json:"id" validate:"required"`
	}

	var reqs []request
	err := ctx.ShouldBindJSON(&reqs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var findedRoles []model.Role

	for _, req := range reqs {
		role, err := s.store.Role().ByID(req.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		} else {
			findedRoles = append(findedRoles, role)
		}
	}

	ctx.JSON(http.StatusOK, findedRoles)
}

func (s *server) DeleteRole(ctx *gin.Context) {
	type request struct {
		Id uint `json:"id" validate:"required"`
	}

	var reqs []request
	err := ctx.ShouldBindJSON(&reqs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, req := range reqs {
		err := s.store.Role().DeleteByID(req.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Запрос на удаление успешно выполнен"})
}

func (s *server) UpdateRole(ctx *gin.Context) {
	var role model.Role
	err := ctx.ShouldBindJSON(&role)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err = s.store.Role().Update(role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, role)
}

// UserRole ...

func (s *server) CreateUserRole(ctx *gin.Context) {
	var reqs []model.UserRole

	err := ctx.ShouldBindJSON(&reqs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var createdUR []model.UserRole
	for _, req := range reqs {
		ur, err := s.store.UserRole().Create(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		} else {
			createdUR = append(createdUR, ur)
		}
	}

	ctx.JSON(http.StatusCreated, createdUR)
}

func (s *server) GetUserRoleById(ctx *gin.Context) {
	type request struct {
		ID int `json:"id"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ur, err := s.store.UserRole().ByID(uint(req.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, ur)
}

func (s *server) UpdateUserRole(ctx *gin.Context) {
	var ur model.UserRole
	err := ctx.ShouldBindJSON(&ur)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ur, err = s.store.UserRole().Update(ur)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, ur)
}

func (s *server) GetUserRoleByRoleId(ctx *gin.Context) {
	type request struct {
		RoleID uint `json:"role_id"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ur, err := s.store.UserRole().ByRoleID(req.RoleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, ur)
}

func (s *server) GetAllUserRole(ctx *gin.Context) {
	ur, err := s.store.UserRole().All()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, ur)
}

func (s *server) GetUserRoleByUserId(ctx *gin.Context) {
	type request struct {
		UserID uint `json:"user_id"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tcs, err := s.store.UserRole().ByUserID(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tcs)
}

func (s *server) GetRoleUser(ctx *gin.Context) {
	type request struct {
		ID uint `json:"id"`
	}

	var req request
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var role model.Role
	err = s.store.Role().UsersByIdRole(req.ID, &role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, role)
}
