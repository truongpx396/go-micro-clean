package http

import (
	"net/http"
	"project/common"
	"project/modules/auth/entity"

	"github.com/gin-gonic/gin"
)

// Auth godoc
// @Summary Authenticate User
// @Description Loggin User By UserName and Password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param UserInfo body usermodel.UserLogin true "Login body"
// @Success 200 {object} entity.TokenResponse
// @Failure 400 {object} common.AppError
// @Failure 401 {object} common.AppError
// @Router /auth/authenticate [post]
func (api *authHandler) LoginHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var data entity.AuthEmailPassword

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		response, err := api.business.Login(c.Request.Context(), &data)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, response)
	}
}
