package database

import (
	"awesomeProject1/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var connectionString = "host=localhost port=5432 user=postgres password=password dbname=go sslmode=disable TimeZone=UTC"

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
	err := DB.AutoMigrate(&models.User{}, &models.Role{})
	roles := []models.Role{
		{Role: "admin"},
		{Role: "user"},
	}
	for _, role := range roles {
		DB.FirstOrCreate(&role, models.Role{Role: role.Role})
	}
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
