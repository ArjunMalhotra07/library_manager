package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := h.Repo.DeleteBook(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Error deleting book", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Successfully deleted book"})
}
