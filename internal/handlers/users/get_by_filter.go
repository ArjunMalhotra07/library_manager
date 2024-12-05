package users

import "github.com/gin-gonic/gin"

func GetUserByFilter(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "Guest")
	age := ctx.Query("age")
	ctx.JSON(200, gin.H{"name": name, "age": age})
}
