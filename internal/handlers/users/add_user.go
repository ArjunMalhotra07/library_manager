package users

import "github.com/gin-gonic/gin"

func AddUser(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	ctx.JSON(201, gin.H{"status": "user created", "user": user})
}
