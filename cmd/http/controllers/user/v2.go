package user

import (
	"github.com/gin-gonic/gin"
	"mono-base/internal/usecases/user"
)

var _ Controller = (*ControllerV2)(nil)

type ControllerV2 struct {
	loginUseCase user.LoginUseCase
}

func NewUserControllerV2(loginUseCase user.LoginUseCase) *ControllerV2 {
	return &ControllerV2{loginUseCase: loginUseCase}
}

// Login
// @Router /v2/login [post]
// @Summary Login v2
// @Description Login v2
// @Tags User v2
// @Accept json
// @Produce json
// @Param payload body user.LoginRequest true "payload"
// @Success 200 {object} rest.BaseResponse
func (u *ControllerV2) Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "login"})
}
