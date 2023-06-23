package main

import (
	"learn-clean-arch/handlers"
	"learn-clean-arch/models"
	"learn-clean-arch/repository"
	"learn-clean-arch/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_buku?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection error")
	}

	db.AutoMigrate(&models.Book{})

	bookRepository := repository.NewBookRepository(db)
	bookService := services.NewBookService(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)

	router := gin.Default()

	api := router.Group("/api")
	api.GET("/books", bookHandler.FindAll)
	api.POST("/books", bookHandler.Create)
	api.GET("/books/:id", bookHandler.FindByID)
	api.PUT("/books/:id", bookHandler.Update)
	api.DELETE("/books/:id", bookHandler.Delete)

	router.Run()
}
