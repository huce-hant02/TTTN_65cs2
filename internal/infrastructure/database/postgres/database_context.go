package postgres

import (
	_ "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"sync"
)

var lock sync.Mutex

type DBConfig struct {
	Driver   string `json:"driver,omitempty" yaml:"driver"`
	User     string `json:"user,omitempty" yaml:"user"`
	Password string `json:"password,omitempty" yaml:"password"`
	Protocol string `json:"protocol,omitempty" yaml:"protocol"`
	Host     string `json:"host,omitempty" yaml:"host"`
	Port     string `json:"port,omitempty" yaml:"port"`
	Schema   string `json:"schema,omitempty" yaml:"schema"`
}

var instance *sqlx.DB

func getDbConfig() *DBConfig {
	dbConfig := &DBConfig{}
	if err := viper.UnmarshalKey("database", dbConfig); err != nil {
		panic(err)
	}
	return dbConfig
}

func GetDBContext() *sqlx.DB {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			dbConfig := getDbConfig()
			dbSource := fmt.Sprintf(
				"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
				dbConfig.User,
				dbConfig.Password,
				dbConfig.Host,
				dbConfig.Port,
				dbConfig.Schema,
			)
			db, err := sqlx.Open(dbConfig.Driver, dbSource)
			if err != nil {
				panic(err)
			}
			db.SetMaxIdleConns(10)
			db.SetMaxOpenConns(20)
			stats := db.Stats()
			_ = stats
			instance = db
		}
	}
	return instance
}
