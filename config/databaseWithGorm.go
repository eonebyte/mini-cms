package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBGorm *gorm.DB

func ConnectDB() {
	host := "localhost"
	port := "3306"
	dbname := "myweb"
	username := "root"
	password := ""

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"
	var err error
	DBGorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("Cannot connect database")
	}
	DBGorm.AutoMigrate()
}
