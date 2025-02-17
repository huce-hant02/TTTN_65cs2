package services

import (
	"github.com/google/wire"
	"mono-base/internal/services/user"
)

var UserServiceProvider = wire.NewSet(user.NewAuthService)
