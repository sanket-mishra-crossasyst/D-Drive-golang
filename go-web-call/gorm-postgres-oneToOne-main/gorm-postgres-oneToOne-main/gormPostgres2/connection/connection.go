package connection

import (
	"demo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var dsn = "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable"

func DatabaseConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if db != nil {
		log.Println("database connected")
	}
	if err != nil {
		log.Println(err)
	}

	db.AutoMigrate(models.Employee{}, models.Department{})
	return db

}
