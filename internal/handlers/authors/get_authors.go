package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AuthorHandler) GetAuthors(c *gin.Context) {
	authors, err := h.Repo.GetAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failure", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Success", "data": authors})
}
