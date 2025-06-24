package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"task_cart/config"
)

var db *gorm.DB

const dbUrl = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func newDB(cfg *config.DbConf) *gorm.DB {
	const op = "pkg.db.newDB"
	connectionString := fmt.Sprintf(dbUrl, cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)

	connection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
		//log.ErrorLog.Panicf("%s: %s", op, err.Error())
	}

	return connection
}

func MustStartDB(cfg *config.DbConf) *gorm.DB {
	if db == nil {
		db = newDB(cfg)
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
		//log.ErrorLog.Panicf("%s: %s", op, err.Error())
	}

	if err = sqlDB.Ping(); err != nil {
		panic(err)
		//log.ErrorLog.Panicf("%s: %s", op, err.Error())
	}

	return db
}

func MustCloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
		//log.ErrorLog.Panicf("%s: %s", op, err.Error())
	}

	if err = sqlDB.Close(); err != nil {
		panic(err)
		//log.ErrorLog.Panicf("%s: can't close connection: %s", op, err.Error())
	}
}
