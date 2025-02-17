package user

import (
	"github.com/gin-gonic/gin"
	"mono-base/internal/usecases/user"
)

var _ Controller = (*ControllerV1)(nil)

type ControllerV1 struct {
	loginUseCase user.LoginUseCase
}

func NewUserControllerV1(loginUseCase user.LoginUseCase) *ControllerV1 {
	return &ControllerV1{loginUseCase: loginUseCase}
}

// Login
// @Router /v1/login [post]
// @Summary Login
// @Description Login
// @Tags User
// @Accept json
// @Produce json
// @Param payload body user.LoginRequest true "payload"
// @Success 200 {object} rest.BaseResponse
func (u *ControllerV1) Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "login"})
}
