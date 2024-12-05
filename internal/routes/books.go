package routes

import (
	handlers "github.com/ArjunMalhotra07/internal/handlers/books"
	repo "github.com/ArjunMalhotra07/internal/repo/books"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BooksAPIs(router *gin.RouterGroup, driver *gorm.DB) {
	bookRepo := repo.NewGormBookRepo(driver)
	bookHandler := handlers.NewBookHandler(bookRepo)
	router.GET("/", bookHandler.GetAllBooks)
	router.POST("/", bookHandler.AddBook)
	router.PUT("/:id", bookHandler.EditBook)
	router.DELETE("/:id", bookHandler.DeleteBook)
}
