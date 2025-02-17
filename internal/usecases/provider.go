package usecases

import (
	"github.com/google/wire"
	edit_history "mono-base/internal/usecases/edit-history"
	"mono-base/internal/usecases/user"
)

var UserUseCaseProviders = wire.NewSet(
	user.NewGetUserByIdUseCase,
	user.NewLoginUseCase,
	edit_history.NewGetInfoInputUseCase,
)
