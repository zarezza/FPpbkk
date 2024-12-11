package config

import (
	"log"
	"os"

	"final-project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database error:", err)
	}
}

func Migrate(models ...interface{}) {
	err := DB.AutoMigrate(models...)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

func Seed() {
	books := []models.Book{
		{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Publisher: "Scribner", ISBN: "9780743273565", Year: 1925, Category: "Fiction", UserID: 1},
		{Title: "To Kill a Mockingbird", Author: "Harper Lee", Publisher: "J.B. Lippincott & Co.", ISBN: "9780061120084", Year: 1960, Category: "Fiction", UserID: 2},
		{Title: "1984", Author: "George Orwell", Publisher: "Secker & Warburg", ISBN: "9780451524935", Year: 1949, Category: "Dystopian", UserID: 1},
		{Title: "Dune", Author: "Frank Herbert", Publisher: "KPG", ISBN: "9786231340061", Year: 1949, Category: "Dystopian", UserID: 2},
	}

	for _, book := range books {
		DB.Create(&book)
	}
}

func DropTables(models ...interface{}) {
	err := DB.Migrator().DropTable(models...)
	if err != nil {
		log.Fatal("Failed to drop tables:", err)
	}
}
