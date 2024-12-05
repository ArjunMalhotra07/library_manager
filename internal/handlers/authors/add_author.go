package handlers

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/ArjunMalhotra07/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *AuthorHandler) AddAuthor(c *gin.Context) {
	var author model.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failure", "error": "Error binding JSON"})
		return
	}

	authorID, err := exec.Command("uuidgen").Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failure", "error": "Error generating user ID"})
		return
	}
	author.ID = strings.TrimSpace(string(authorID))

	if err := h.Repo.AddAuthor(&author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failure", "error": "Error creating author"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Success", "data": author})
}
