package handlers

import (
	"net/http"

	"github.com/ArjunMalhotra07/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *AuthorHandler) EditAuthor(c *gin.Context) {
	id := c.Param("id")
	var updatedAuthor model.Author
	if err := c.ShouldBindJSON(&updatedAuthor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failure", "error": "Error binding JSON"})
		return
	}

	author, err := h.Repo.EditAuthor(id, &updatedAuthor)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Failure", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Success", "data": author})
}
