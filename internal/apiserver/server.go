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
	//	s.router.Use(cors.New(cors.DefaultConfig()))

	/*	userGroup := s.router.Group("/user")
		{
			userGroup.POST("/signin")
			userGroup.POST("/signout")
			userGroup.POST("/update")
		}
	*/
	usersGroup := s.router.Group("/users")
	{
		usersGroup.POST("", s.CreateUser)
	}
}

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
	ctx.JSON(http.StatusOK, user)
}
