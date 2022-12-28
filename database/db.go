package database

import (
	"fmt"
	"log"
	"wiaoj/authorization/config"
	"wiaoj/authorization/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// Connect to database
	dsn := config.GetPostgresDSN()
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Exec("DROP DATABASE IF EXISTS users;").Error; err != nil {
		panic(err)
	}

	if err := DB.Exec("CREATE DATABASE users").Error; err != nil {
		panic(err)
	}

	dsn = config.GetPostgresDSN()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate tables
	DB.AutoMigrate(&models.User{}, &models.PhoneNumber{}, &models.Address{}, &models.UserClaim{})

	// Create Super User
	var user = &models.User{
		Email: "admin@admin.com",
		Hash:  "08f46d939d7ff7ebb1df9de1cc246135c7a8694cabf0e1c1cc67c50e12832f08",
	}

	if result := DB.Create(&user); result.Error != nil {
		fmt.Println("Couldn't create super user")
	}

	var admin = &models.UserClaim{
		UserID:      user.ID,
		AccessLevel: 1,
	}

	if result := DB.Create(&admin); result.Error != nil {
		fmt.Println("Couldn't create super user")
	}

	return DB
}
