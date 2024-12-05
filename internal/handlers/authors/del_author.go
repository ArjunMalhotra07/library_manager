package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	if err := h.Repo.DeleteAuthor(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Failure", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Successfully deleted author"})
}
