package authors

import (
	"net/http"

	"github.com/ArjunMalhotra07/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EditAuthor(c *gin.Context, driver *gorm.DB) {
	authorID := c.Param("id")
	var inputAuthor model.Author
	if err := c.ShouldBindJSON(&inputAuthor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failure", "error": "Error binding json"})
		return
	}
	var existingAuthor model.Author
	if err := driver.Where("id = ?", authorID).Preload("Books").First(&existingAuthor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Failure", "error": err.Error()})
		return
	}
	existingAuthor.Name = inputAuthor.Name
	if err := driver.Save(&existingAuthor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failure", "error": "Error saving author"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Success", "error": "Author updated", "data": existingAuthor})

}
