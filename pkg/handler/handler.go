package handler

import (
	"first-rest-api/pkg/service"

	"github.com/gin-gonic/gin"
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
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			itemsInList := lists.Group("/:id/items")
			{
				itemsInList.POST("/", h.createItem)
				itemsInList.GET("/", h.getAllItems)
			}
		}
		items := api.Group("/items")
		{
			items.GET("/:id", h.getItemById)   // получить item по ID
			items.PUT("/:id", h.updateItem)    // обновить item
			items.DELETE("/:id", h.deleteItem) // удалить item
		}
	}

	return router
}
