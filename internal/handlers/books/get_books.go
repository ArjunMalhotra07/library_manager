package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.Repo.GetAllBooks()
	fmt.Println("here")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failure", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Success", "data": books})
}
