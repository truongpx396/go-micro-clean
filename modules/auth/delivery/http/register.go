package http

import (
	"net/http"
	"project/common"
	"project/modules/auth/entity"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Register New User
// @Description Create a new user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param RegisterData body entity.AuthRegister true "Auth data"
// @Success 200 {object} entity.SuccessResponse
// @Failure 400 {object} common.AppError
// @Failure 401 {object} common.AppError
// @Router /auth/register [post]
func (api *authHandler) RegisterHdl() func(*gin.Context) {
	return func(c *gin.Context) {
		var data entity.AuthRegister

		if err := c.ShouldBind(&data); err != nil {
			common.WriteErrorResponse(c, common.ErrInvalidRequest(err))
			return
		}

		err := api.business.Register(c.Request.Context(), &data)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, entity.SuccessResponse{Data: true})
	}
}
