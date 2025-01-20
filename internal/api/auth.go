package api

import (
	"go-micro-clean/pkg/rpcclient"
	"go-micro-clean/proto/auth"
	"go-micro-clean/tools/a2r"

	"github.com/gin-gonic/gin"
)

type AuthApi rpcclient.Auth

func NewAuthApi(client rpcclient.Auth) AuthApi {
	return AuthApi(client)
}

// @Summary User login
// @Description Get the user token
// @Tags Authentication
// @ID UserToken
// @Accept json
// @Param OperationId header string true "OperationId"
// @Param req body auth.LoginRequest true "Email and Password"
// @Produce json
// @Success 200 {object} auth.LoginResponse
// @Failure 500 {object} error "ERRCODE is 500 generally an internal error of the server"
// @Failure 400 {object} error "Errcode is 400, which is generally a parameter input error."
// @Router /auth/login [post]
func (o *AuthApi) Login(c *gin.Context) {
	a2r.Call(auth.AuthServiceClient.Login, o.Client, c)
}

// @Summary Register new user
// @Description Register a new user
// @Tags Authentication
// @ID RegisterUser
// @Accept json
// @Param OperationId header string true "OperationId"
// @Param req body auth.RegisterRequest true "User registration data"
// @Produce json
// @Success 200 {object} auth.RegisterResponse
// @Failure 500 {object} error "ERRCODE is 500 generally an internal error of the server"
// @Failure 400 {object} error "Errcode is 400, which is generally a parameter input error."
// @Router /auth/register [post]
func (o *AuthApi) RegisterUser(c *gin.Context) {
	a2r.Call(auth.AuthServiceClient.Register, o.Client, c)
}

// @Summary Introspect token
// @Description Introspect the given token
// @Tags Authentication
// @ID IntrospectToken
// @Accept json
// @Param OperationId header string true "OperationId"
// @Param req body entity.IntrospectReq true "Token introspection request"
// @Produce json
// @Success 200 {object} entity.IntrospectResp
// @Failure 500 {object} error "ERRCODE is 500 generally an internal error of the server"
// @Failure 400 {object} error "Errcode is 400, which is generally a parameter input error."
// @Router /auth/introspect_token [post]
func (o *AuthApi) IntrospectToken(c *gin.Context) {
	a2r.Call(auth.AuthServiceClient.IntrospectToken, o.Client, c)
}
