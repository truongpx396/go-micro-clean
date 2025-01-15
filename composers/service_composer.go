package composers

import (
	"context"
	"project/common"

	authApi "project/internal/auth/delivery/http"
	authRepo "project/internal/auth/repository/postgre"
	authUserRpc "project/internal/auth/repository/rpc-client"
	authUsecase "project/internal/auth/usecase"

	itemApi "project/internal/item/delivery/http"
	itemRepo "project/internal/item/repository/postgre"
	itemUsecase "project/internal/item/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthService interface {
	LoginHdl() func(*gin.Context)
	RegisterHdl() func(*gin.Context)
}

func ComposeAuthAPIService(ctx context.Context, db *gorm.DB) AuthService {
	jwtComp := common.NewJWT("jwt")

	authRepo := authRepo.NewPostgreRepository(db)
	hasher := new(common.Hasher)

	userRepo := authUserRpc.NewClient(composeUserRPCClient(ctx))
	biz := authUsecase.NewAuthUsecase(authRepo, userRepo, jwtComp, hasher)
	serviceAPI := authApi.NewAuthHandler(ctx, biz)

	return serviceAPI
}

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
