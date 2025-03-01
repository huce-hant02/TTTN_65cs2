// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"mono-base/internal/infrastructure/database/postgres"
)

// Injectors from wire.go:

func wireApp(app *App) error {
	migration := postgres.NewPSQLMigration()
	error2 := inject(app, migration)
	return error2
}
