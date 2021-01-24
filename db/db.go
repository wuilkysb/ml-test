package database

import (
	"ml-mutant-test/config"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/gommon/log"
	"sync"
)

var instance *pg.DB
var once sync.Once

func ConnInstance() *pg.DB {
	once.Do(func() {
		instance = getConnection()
	})

	return instance
}

func getConnection() *pg.DB {
	address := fmt.Sprintf("%s:%d", config.Environments().PgHost, config.Environments().PgPort)

	db := pg.Connect(&pg.Options{
		Addr:         address,
		User:         config.Environments().PgUser,
		Password:     config.Environments().PgPassword,
		Database:     config.Environments().PgName,
		ReadTimeout:  config.Environments().PgTimeout,
		WriteTimeout: config.Environments().PgTimeout,
		DialTimeout:  config.Environments().PgTimeout,
		PoolSize:     config.Environments().PgPoolSize,
	})

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		log.Error(err.Error())
	}

	log.Info("Database successfully connected")

	return db
}
