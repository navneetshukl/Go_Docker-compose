package database

import (
	"docker-compose/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
	dsn := "host=db port=5432 dbname=docker-compose user=postgres password=password"

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database ", err)
		return nil, err
	}
	fmt.Println("Connected to Database")
	return DB, nil
}
func MigrateDatabase() {
	DB, err := ConnectToDatabase()
	if err != nil {
		log.Fatal("There is error connecting to database ", err)
		return
	}
	DB.AutoMigrate(&models.User{})
}

func InsertUser(name, email, phone string) error {
	DB, _ := ConnectToDatabase()
	user := models.User{
		Name:  name,
		Email: email,
		Phone: phone,
	}
	err := DB.Create(&user)
	if err != nil {
		return err.Error
	}
	return nil
}

func FindAllUser() ([]models.User) {
	DB, _ := ConnectToDatabase()
	var users []models.User
	if err := DB.Find(&users).Error; err != nil {
		panic("Failed to fetch data from database")
		return nil
	}
	return users
}
