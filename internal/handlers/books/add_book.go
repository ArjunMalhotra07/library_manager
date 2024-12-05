package handlers

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/ArjunMalhotra07/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *BookHandler) AddBook(c *gin.Context) {
	var book *model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failure", "error": "Error binding JSON"})
		return
	}
	bookID, err := exec.Command("uuidgen").Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failure", "error": "Error generating book ID"})
		return
	}
	book.ID = strings.TrimSpace(string(bookID))
	if err := h.Repo.AddBook(book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failure", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Success", "data": book})

}
