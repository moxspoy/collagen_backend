package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() error {
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/collagen?charset=utf8mb4&parseTime=True&loc=Local"
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
	err = MigrateAppConfig()
	if err != nil {
		log.Fatal("Migration App Config Error:" + err.Error())
	}
}
