package routes

import (
	"project/composers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	itemService := composers.ComposeItemAPIServie(db)

	api := r.Group("/api/v1")
	{
		api.POST("/items", itemService.CreateItem)
		api.GET("/items", itemService.ListItems)
		api.GET("/items/:id", itemService.GetItemByID)
		api.PUT("/items/:id", itemService.UpdateItem)
		api.PATCH("/items/:id", itemService.PatchItem)
		api.DELETE("/items/:id", itemService.DeleteItem)
	}
}
