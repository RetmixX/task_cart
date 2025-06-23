package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"task_cart/config"
)

var db *gorm.DB

const dbUrl = "host=%s user=%s password=%s dbname=%s sslmode=disable"

func newDB(cfg *config.DbConf, log *log.Logger) *gorm.DB {
	const op = "pkg.db.newDB"
	connectionString := fmt.Sprintf(dbUrl, cfg.Host, cfg.User, cfg.Password, cfg.Name)

	connection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
		//log.ErrorLog.Panicf("%s: %s", op, err.Error())
	}

	return connection
}

func MustStartDB(cfg *config.DbConf, log *log.Logger) *gorm.DB {
	const op = "pkg.db.MustStartDB"
	if db == nil {
		db = newDB(cfg, log)
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

	//log.InfoLog.Printf("%s: Connect to db\n", op)

	return db
}

func MustCloseDB(db *gorm.DB, log *log.Logger) {
	const op = "pkg.db.MustCloseDB"

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
		//log.ErrorLog.Panicf("%s: %s", op, err.Error())
	}

	if err = sqlDB.Close(); err != nil {
		panic(err)
		//log.ErrorLog.Panicf("%s: can't close connection: %s", op, err.Error())
	}

	//log.InfoLog.Printf("%s: Success close connection\n", op)
}
