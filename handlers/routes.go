package handlers

import (
	"github.com/ArjunMalhotra07/handlers/authors"
	"github.com/ArjunMalhotra07/handlers/books"
	"github.com/ArjunMalhotra07/handlers/borrows"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OpenServer(gormDriver *gorm.DB) {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "Hey"})
	})

	var authorAPIs *gin.RouterGroup = r.Group("/author")
	{
		AuthorAPIs(authorAPIs, gormDriver)
	}
	var booksAPIs *gin.RouterGroup = r.Group("/book")
	{
		BooksAPIs(booksAPIs, gormDriver)
	}
	var borrowAPIs *gin.RouterGroup = r.Group("/borrow")
	{
		BorrowAPIs(borrowAPIs, gormDriver)
	}

	r.Run(":8080")
}
func AuthorAPIs(router *gin.RouterGroup, driver *gorm.DB) {
	router.GET("/", func(ctx *gin.Context) {
		authors.GetAuthors(ctx, driver)
	})
	router.POST("/", func(ctx *gin.Context) {
		authors.AddAuthor(ctx, driver)
	})
	router.PUT("/:id", func(ctx *gin.Context) {
		authors.EditAuthor(ctx, driver)
	})
	router.DELETE("/:id", func(ctx *gin.Context) {
		authors.DeleteAuthor(ctx, driver)
	})
}

func BooksAPIs(router *gin.RouterGroup, driver *gorm.DB) {
	router.GET("/", func(ctx *gin.Context) {
		books.GetAllBooks(ctx, driver)
	})
	router.POST("/", func(ctx *gin.Context) {
		books.AddBook(ctx, driver)
	})
	router.PUT("/", func(ctx *gin.Context) {
		books.EditBook(ctx, driver)
	})
	router.DELETE("/", func(ctx *gin.Context) {
		books.DeleteBook(ctx, driver)
	})
}

func BorrowAPIs(router *gin.RouterGroup, driver *gorm.DB) {
	router.GET("/getborrows", func(ctx *gin.Context) {
		borrows.GetAllBorrows(ctx, driver)
	})
	router.POST("/borrowbook", func(ctx *gin.Context) {
		borrows.BorrowBook(ctx, driver)
	})
	router.DELETE("/returnbook", func(ctx *gin.Context) {
		borrows.ReturnBook(ctx, driver)
	})
}
