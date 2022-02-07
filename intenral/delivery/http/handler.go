package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lov3allmy/todo-app/intenral/service"
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

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.readAllLists)
			lists.GET("/:list_id", h.readListByID)
			lists.PUT("/:list_id", h.updateList)
			lists.DELETE("/:list_id", h.deleteList)

			items := lists.Group(":list_id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.readAllItems)
			}
		}

		items := api.Group("/items")
		{
			items.GET("/:item_id", h.readItemByID)
			items.PUT("/:item_id", h.updateItem)
			items.DELETE("/:item_id", h.deleteItem)
		}
	}

	return router
}
