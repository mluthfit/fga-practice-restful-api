package config

import (
	"fga-practice-rest-api/structs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "127.0.0.1"
	port     = "3306"
	username = "root"
	password = ""
	dbName   = "go-practice-rest-api"
)

func DBInit() *gorm.DB {
	var dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbName,
	)

	var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.Debug().AutoMigrate(structs.Person{})

	return db
}
