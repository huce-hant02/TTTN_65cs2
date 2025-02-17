package database

import (
	"github.com/google/wire"
	"mono-base/internal/infrastructure/database/postgres"
)

var DBProvider = wire.NewSet(
	postgres.GetDBContext,
	postgres.NewPSQLMigration,
	postgres.NewUserRepository,
	postgres.NewEditHistoryRepository,
)
