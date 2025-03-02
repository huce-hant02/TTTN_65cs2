//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"mono-base/cmd/http/controllers"
	"mono-base/internal/infrastructure/database"
	"mono-base/internal/services"
	"mono-base/internal/usecases"
)

func wireApp(app *App) error {
	wire.Build(
		database.DBProvider,
		controllers.ControllerProviders,
		usecases.UserUseCaseProviders,
		services.UserServiceProvider,
		inject,
	)
	return nil
}
