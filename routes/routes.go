package routes

import (
	"project/modules/item/delivery/http"
	repository "project/modules/item/repository/postgre"
	"project/modules/item/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	itemRepo := repository.NewItemRepository(db)
	itemUC := usecase.NewItemUsecase(itemRepo)
	itemHandler := http.NewItemHandler(itemUC)

	api := r.Group("/api/v1")
	{
		api.POST("/items", itemHandler.CreateItem)
		api.GET("/items", itemHandler.ListItems)
		api.GET("/items/:id", itemHandler.GetItemByID)
		api.PUT("/items/:id", itemHandler.UpdateItem)
		api.PATCH("/items/:id", itemHandler.PatchItem)
		api.DELETE("/items/:id", itemHandler.DeleteItem)
	}
}
