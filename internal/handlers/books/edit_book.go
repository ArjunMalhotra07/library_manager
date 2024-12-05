package handlers

import (
	"net/http"

	"github.com/ArjunMalhotra07/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *BookHandler) EditBook(c *gin.Context) {
	id := c.Param("id")
	var book *model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failure", "error": "Error binding JSON"})
		return
	}
	book, err := h.Repo.EditBook(id, book)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Error editing book", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Success", "data": book})
}
