package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() error {
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/flop?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = database
	return nil
}

func SetupDatabase() {
	err := ConnectDatabase()
	if err != nil {
		return
	}

	err = RunAutoMigration(DB)
	if err != nil {
		log.Fatal("Migration Error:" + err.Error())
	}
}
