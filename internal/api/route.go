package api

import (
	"context"
	"go-micro-clean/common/config"
	"go-micro-clean/pkg/rpcclient"
	"go-micro-clean/tools/log"
	"go-micro-clean/tools/mw"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// swag init -g cmd/main.go -o docs
// swag init --parseInternal --pd --dir cmd/api,internal/api/ -g ../../internal/api/route.go --output cmd/api/docs

// @title go-micro-clean API
// @version 1.0
// @description  go-micro-clean API server document

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:10002
// @BasePath /

//	@schemes http https

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						token
//	@description				Description for what is this security definition being used

//// @securityDefinitions.basic	OperationId
//// @in							header
//// @name						OperationId
//// @description				OperationId for link tracking

func NewGinRouter(ctx context.Context) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// Use CORS middleware
	r.Use(gin.Recovery(), mw.CorsHandler())

	log.Info1("load config", "config", config.Config)
	// r.Use(gin.Recovery(), mw.CorsHandler(), mw.GinParseOperationID())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ParseToken := GinParseToken(ctx)

	authRpc := rpcclient.NewAuth(ctx)

	// certificate
	authRouterGroup := r.Group("/auth")
	{
		a := NewAuthApi(*authRpc)
		authRouterGroup.POST("/login", a.Login)
		authRouterGroup.POST("/register", a.RegisterUser)
		authRouterGroup.POST("/intropect_token", ParseToken, a.IntrospectToken)
	}

	// init rpc client here
	// userRpc := rpcclient.NewUser(discov)

	// u := NewUserApi(*userRpc)
	// m := NewMessageApi(messageRpc, userRpc)

	// userRouterGroup := r.Group("/user")
	// {
	// 	userRouterGroup.POST("/user_register", u.UserRegister)
	// 	userRouterGroup.POST("/update_user_info", ParseToken, u.UpdateUserInfo)
	// 	userRouterGroup.POST("/get_users_info", ParseToken, u.GetUsersPublicInfo)
	// 	userRouterGroup.POST("/get_users", ParseToken, u.GetUsers)
	// }
	return r
}

func GinParseToken(ctx context.Context) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()
	}
}
