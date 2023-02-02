package databaseconnection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"onetomany/model"
)

var dsn = "host=localhost user=postgres password=root dbname=oneToMany port=5432 sslmode=disable"

func Connection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if db != nil {
		log.Println("database connected successfully")
	}
	if err != nil {
		log.Println("database connection failed")
	}
	db.AutoMigrate(model.Employee{}, model.Department{})
	return db

}
