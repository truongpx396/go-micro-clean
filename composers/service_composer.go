package composers

import (
	itemApi "go-micro-clean/internal/item/delivery/http"
	itemRepo "go-micro-clean/internal/item/repository/postgre"
	itemUsecase "go-micro-clean/internal/item/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthService interface {
	LoginHdl() func(*gin.Context)
	RegisterHdl() func(*gin.Context)
}

// func ComposeAuthAPIService(ctx context.Context, db *gorm.DB) AuthService {
// 	jwtComp := common.NewJWT("jwt")

// 	authRepo := authRepo.NewPostgreRepository(db)
// 	hasher := new(common.Hasher)

// 	userRepo := authUserRpc.NewClient(composeUserRPCClient(ctx))
// 	biz := authUsecase.NewAuthUsecase(authRepo, userRepo, jwtComp, hasher)
// 	serviceAPI := authApi.NewAuthHandler(ctx, biz)

// 	return serviceAPI
// }

type ItemAPIService interface {
	CreateItem(c *gin.Context)
	ListItems(c *gin.Context)
	GetItemByID(c *gin.Context)
	UpdateItem(c *gin.Context)
	PatchItem(c *gin.Context)
	DeleteItem(c *gin.Context)
}

func ComposeItemAPIServie(db *gorm.DB) ItemAPIService {
	itemRepo := itemRepo.NewItemRepository(db)
	itemUC := itemUsecase.NewItemUsecase(itemRepo)
	return itemApi.NewItemHandler(itemUC)
}
