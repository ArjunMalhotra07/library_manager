package routes

import (
	"github.com/ArjunMalhotra07/internal/handlers/borrows"
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
