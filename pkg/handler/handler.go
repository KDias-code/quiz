package handler

import (
	"github.com/KDias-code/pkg/service"
	"github.com/gin-gonic/gin"
	// "github.com/go-delve/delve/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	user := router.Group("/auth")
	{
		user.POST("/sign-up", h.SignUp)
		user.POST("/sign-in", h.SignIn)
	}

	//game := router.Group("/game", h.userIdentity)
	//{
	//	start := router.Group("/start")
	//	{
	//		start.POST("/:id", h.quiz)
	//	}
	//}

	return router
}
