package http

import (
	"context"
	"net/http"
	"go-micro-clean/common"
	"go-micro-clean/internal/user/entity"

	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	GetUserProfile(ctx context.Context) (*entity.User, error)
}

type userApi struct {
	business UserUsecase
}

func NewAPI(business UserUsecase) *userApi {
	return &userApi{business: business}
}

// Profile godoc
// @Summary User profile
// @Description Get profile of user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} entity.User
// @Failure 400 {object} common.AppError
// @Failure 401 {object} common.AppError
// @Security	ApiKeyAuth
// @Router /user/profile [get]
func (api *userApi) GetUserProfileHdl() func(c *gin.Context) {
	return func(c *gin.Context) {

		//requester := c.MustGet(core.KeyRequester).(core.Requester)
		//ctx := core.ContextWithRequester(c.Request.Context(), requester)

		user, err := api.business.GetUserProfile(c.Request.Context())

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		//user.Mask(common.MaskTypeUser)

		c.JSON(http.StatusOK, user)
	}
}
