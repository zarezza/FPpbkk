package main

import (
	"final-project/config"
	controllers "final-project/controllers/BookController"
	"final-project/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.DBConnection()
	config.DropTables(&models.Book{})
	config.Migrate(&models.Book{})
	config.Seed()

	bookModel := &models.BookModel{DB: config.DB}
	bookController := &controllers.BookController{Model: bookModel}

	r := gin.Default()
	r.LoadHTMLGlob("views/book/*")
	r.Static("/css", "./views/css")

	r.GET("/books", bookController.Index)
	r.GET("/books/add", bookController.Add)
	r.POST("/books", bookController.Add)
	r.GET("/books/edit/:id", bookController.Edit)
	r.POST("/books/:id", bookController.Edit)
	r.GET("/books/delete/:id", bookController.Delete)

	r.Run()
}
