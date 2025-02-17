package postgres

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/viper"
	"mono-base/internal/repositories/migration"
	"os"

	_ "database/sql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const migrationDir = "internal/infrastructure/database/sql_migration"

type PSQLMigration struct {
	Driver   string
	DBSource string
}

func NewPSQLMigration() migration.Migration {
	var dbConfig *DBConfig
	err := viper.UnmarshalKey("database", &dbConfig)
	if err != nil {
		panic(err)
	}
	return &PSQLMigration{
		DBSource: fmt.Sprintf(
			"%s:%s@%s:%s/%s?sslmode=disable",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Schema,
		),
		Driver: dbConfig.Driver,
	}
}

func (m *PSQLMigration) Migrate() {
	log.Info("Start migrate database")

	_, migrationErr := os.Stat(migrationDir)
	if migrationErr != nil {
		if os.IsNotExist(migrationErr) {
			fmt.Printf("cmd/migrations not found: %s \n", migrationErr)
			return
		}
		log.Error("Failed while check status migration")
	}
	migrator, migratorErr := migrate.New(
		fmt.Sprintf("file://%s", migrationDir),
		fmt.Sprintf("%s://%s", m.Driver, m.DBSource),
	)
	if migratorErr != nil {
		panic(migratorErr)
	}
	if upErr := migrator.Up(); upErr != nil {
		noChange := fmt.Sprintf("%s", upErr)
		if noChange != "no change" {
			panic(upErr)
		}
	}
	log.Infof("Migration has been done successfully")
}
