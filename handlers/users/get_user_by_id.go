package users

import (
	"github.com/gin-gonic/gin"
)

func GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, gin.H{"id": id, "name": "Arjun", "email": "arjun7.malhotra@gmail.com"})
}
