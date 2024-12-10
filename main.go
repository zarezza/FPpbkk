package main

import (
	"final-project/config"
	bookControllers "final-project/controllers/BookController"
	userControllers "final-project/controllers/UserController"
	"final-project/middleware"
	"final-project/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database setup
	config.DBConnection()
	config.DropTables(&models.Book{}, &models.User{})
	config.Migrate(&models.Book{}, &models.User{})
	config.Seed()

	// Initialize models and controllers
	bookModel := &models.BookModel{DB: config.DB}
	userModel := &models.UserModel{DB: config.DB}
	bookController := &bookControllers.BookController{Model: bookModel}
	userController := &userControllers.UserController{Model: userModel}

	// Setup Gin router
	r := gin.Default()
	r.LoadHTMLGlob("views/book/*")
	r.Static("/css", "./views/css")

	// Public routes (no authentication required)
	r.GET("/register", userController.Register)
	r.POST("/registers", userController.Register)
	r.GET("/login", userController.Login)
	r.POST("/logins", userController.Login)
	r.GET("/logout", userController.Logout)

	// Protected routes (authentication required)
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		// Book routes
		authorized.GET("/books", bookController.Index)
		authorized.GET("/books/add", bookController.Add)
		authorized.POST("/books", bookController.Add)
		authorized.GET("/books/edit/:id", bookController.Edit)
		authorized.POST("/books/:id", bookController.Edit)
		authorized.GET("/books/delete/:id", bookController.Delete)
	}

	// Default route redirect to login
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/login")
	})

	// Start the server
	r.Run()
}
