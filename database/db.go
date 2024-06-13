package database

import (
	"awesomeProject1/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var connectionString = "host=localhost port=5432 user=postgres password=postgres dbname=godatabse sslmode=disable TimeZone=UTC"

var DB *gorm.DB

func init() {
	Connect()
}

func Connect() {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	DB = db

	Migrate()
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	err = sqlDB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
