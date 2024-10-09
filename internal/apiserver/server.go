package apiserver

import (
	"errors"
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
	userGroup := s.router.Group("/user")
	{
		userGroup.POST("/signout", s.SignOutUserById)
		userGroup.POST("/update", s.UpdateUser)
		userGroup.POST("/pass", s.UpdatePassword)
	}

	usersGroup := s.router.Group("/users")
	{
		usersGroup.POST("", s.CreateUser)
		usersGroup.POST("/signin", s.SignIn)
	}

	teamGroup := s.router.Group("/team")
	{
		teamGroup.GET("", s.GetTeamById)
		teamGroup.POST("/delete", s.DeteleTeamById)
		teamGroup.POST("/update", s.UpdateTeam)

	}

	teamsGroup := s.router.Group("/teams")
	{
		teamsGroup.POST("", s.CreateTeam)
		teamsGroup.GET("", s.GetTeamAll)
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	err := ctx.ShouldBind(&teams)
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
	err := ctx.ShouldBind(&reqs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var teams []model.Team

	for _, req := range reqs {
		team, err := s.store.Team().GetById(req.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue
		} else {
			teams = append(teams, team)
		}
	}

	ctx.JSON(http.StatusOK, teams)
}

func (s *server) DeteleTeamById(ctx *gin.Context) {
	type request struct {
		Id int `json:"id" validate:"required"`
	}

	var reqs []request
	err := ctx.ShouldBind(&reqs)
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
