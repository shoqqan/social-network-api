package handler

import (
	"github.com/gin-gonic/gin"
	"social-network-api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		posts := api.Group("/posts")
		{
			posts.POST("/", h.createPost)
			posts.GET("/", h.getPosts)
			posts.GET("/:post_id", h.getPostById)
			posts.PATCH("/:post_id", h.updatePost)
			posts.DELETE("/:post_id", h.deletePost)
		}

		users := api.Group("/users")
		{
			users.GET("/:user_id/posts", h.getUsersPosts)
		}
	}
	return router
}
