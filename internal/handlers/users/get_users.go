package users

import (
	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	users := []User{
		{Name: "Arjun", Email: "arjun@example.com"},
		{Name: "Nishith", Email: "nishith@example.com"},
		{Name: "Vansh", Email: "vansh@example.com"},
		{Name: "Utkarsh", Email: "utkarsh@example.com"},
	}
	ctx.JSON(200, gin.H{"users": users})
	// ctx.XML(200, gin.H{"message": "Success"})
}
