package api

import (
	"go-micro-clean/pkg/rpcclient"
	"go-micro-clean/proto/auth"
	"go-micro-clean/tools/a2r"

	"github.com/gin-gonic/gin"
)

type authApi rpcclient.Auth

func NewAuthApi(client rpcclient.Auth) authApi {
	return authApi(client)
}

// Auth godoc
// @Summary Authenticate User
// @Description Loggin User By UserName and Password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param UserInfo body auth.LoginRequest true "Email and Password"
// @Success 200 {object} auth.LoginResponse
// @Failure 400 {object} common.AppError
// @Failure 401 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /auth/login [post]
func (o *authApi) Login(c *gin.Context) {
	a2r.Call(auth.AuthServiceClient.Login, o.Client, c)
}

// Register godoc
// @Summary Register New User
// @Description Create a new user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param RegisterData body auth.RegisterRequest true "User registration data"
// @Success 200 {object} auth.RegisterResponse
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /auth/register [post]
func (o *authApi) RegisterUser(c *gin.Context) {
	a2r.Call(auth.AuthServiceClient.Register, o.Client, c)
}

// @Summary Introspect token
// @Description Introspect the given token
// @Tags Authentication
// @Accept json
// @Param OperationId header string true "OperationId"
// @Param req body auth.IntrospectReq true "Token introspection request"
// @Produce json
// @Success 200 {object} auth.IntrospectResp
// @Failure 500 {object} common.AppError
// @Failure 400 {object} common.AppError
// @Router /auth/introspect_token [post]
func (o *authApi) IntrospectToken(c *gin.Context) {
	a2r.Call(auth.AuthServiceClient.IntrospectToken, o.Client, c)
}
