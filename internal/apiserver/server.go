package apiserver

import (
	"errors"
	"html/template"
	"net/http"
	"obstore/internal/model"
	"obstore/internal/store"

	"github.com/gin-gonic/gin"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not autenticated")
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

	s.configureRouter()

	return s
}

func (s *server) configureRouter() {

	s.router.LoadHTMLGlob("frontend/login/*")
	s.router.StaticFS("frontend/login/style.css", http.Dir("/frontend/login"))

	// Маршруты для API
	apiGroup := s.router.Group("/api/v1")
	{
		userGroup := apiGroup.Group("/user")
		{
			userGroup.POST("/signout", s.SignOutUserById)
			userGroup.POST("/update", s.UpdateUser)
			userGroup.POST("/update/pass", s.UpdatePassword)
		}

		usersGroup := apiGroup.Group("/users")
		{
			usersGroup.POST("", s.CreateUser)
			usersGroup.POST("/signin", s.SignIn)
		}

		teamGroup := apiGroup.Group("/team")
		{
			teamGroup.GET("", s.GetTeamById)
			teamGroup.POST("/delete", s.DeteleTeamById)
			teamGroup.POST("/update", s.UpdateTeam)

		}

		teamsGroup := apiGroup.Group("/teams")
		{
			teamsGroup.POST("", s.CreateTeam)
			teamsGroup.GET("", s.GetTeamAll)
		}

		productGroup := apiGroup.Group("/product")
		{
			productGroup.GET("/find/article", s.GetProductByArticle)
			productGroup.GET("/find/strikecode", s.GetProductByStrikeCode)
			productGroup.GET("/find/name", s.GetProductByName)
			productGroup.POST("/update/strikecode", s.UpdateProductStrikeCodeById)
		}

		productsGroup := apiGroup.Group("/products")
		{
			productsGroup.POST("", s.CreateProduct)
			productsGroup.GET("", s.GetProductsAll)
		}

		orderGroup := apiGroup.Group("/order")
		{
			orderGroup.GET("/find/id", s.GetOrderById)
			orderGroup.GET("/find/uid", s.GetOrderByUID)
			orderGroup.GET("/find/num", s.GetOrderByFolioNum)
		}

		ordersGroup := apiGroup.Group("/orders")
		{
			ordersGroup.POST("", s.CreateOrder)
			ordersGroup.GET("", s.GetOrdersAll)
			ordersGroup.GET("/range", s.GetOrdersByDateRange)
			ordersGroup.GET("/driver", s.GetOrdersByDriver)
			ordersGroup.GET("/agent", s.GetOrdersByAgent)
		}

		teamCompositionGroup := apiGroup.Group("/teamcomposition")
		{
			teamCompositionGroup.GET("/", s.GetTeamCompositionById)
			teamCompositionGroup.POST("/update", s.UpdateTeamComposition)
			teamCompositionGroup.GET("/team", s.GetTeamCompositionByTeamId)
			teamCompositionGroup.GET("/user", s.GetTeamCompositionByUserId)

		}

		teamCompositionsGroup := apiGroup.Group("/teamcompositions")
		{
			teamCompositionsGroup.POST("", s.CreateTeamComposition)
			teamCompositionsGroup.GET("", s.GetTeamCompositionAll)

		}
	}

	// Маршруты для сайта
	viewGroup := s.router.Group("/view")
	{

		h1 := func(ctx *gin.Context) {
			tmpl := template.Must(template.ParseFiles("frontend/login/login.html"))

			tmpl.Execute(ctx.Writer, nil)
		}

		viewGroup.GET("/login", h1)

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

	user, err = s.store.User().SignInUser(user.Email, user.Pass)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errIncorrectEmailOrPassword.Error()})
		return
	}

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

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
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
		OrderId int `json:"order_id" validate: "required"`
	}

	var reqs []request
	err := ctx.ShouldBind(&reqs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var orders []model.Order

	for _, req := range reqs {
		order, err := s.store.Order().GetById(req.OrderId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		} else {
			orders = append(orders, order)
		}
	}
}

func (s *server) GetOrderByUID(ctx *gin.Context) {
	type request struct {
		OrderUID int `json:"order_uid" validate: "required"`
	}

	var reqs []request
	err := ctx.ShouldBind(&reqs)
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
}

func (s *server) GetOrderByFolioNum(ctx *gin.Context) {
	type request struct {
		FolioNum int `json:"folio_num" validate: "required"`
	}

	var reqs []request
	err := ctx.ShouldBind(&reqs)
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
	err := ctx.ShouldBindQuery(&req)
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
