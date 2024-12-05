package routes

import (
	handlers "github.com/ArjunMalhotra07/internal/handlers/authors"
	repo "github.com/ArjunMalhotra07/internal/repo/authors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthorAPIs(router *gin.RouterGroup, driver *gorm.DB) {
	authorRepo := repo.NewGormAuthorRepo(driver)
	authorHandler := handlers.NewAuthorHandler(authorRepo)
	router.GET("/", authorHandler.GetAuthors)
	router.POST("/", authorHandler.AddAuthor)
	router.PUT("/:id", authorHandler.EditAuthor)
	router.DELETE("/:id", authorHandler.DeleteAuthor)
}
