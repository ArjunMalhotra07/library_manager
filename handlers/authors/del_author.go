package authors

import (
	"net/http"

	"github.com/ArjunMalhotra07/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteAuthor(c *gin.Context, driver *gorm.DB) {
	authorID := c.Param("id")
	var author model.Author

	if err := driver.Where("id = ? ", authorID).First(&author).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"msg": "Failure", "error": "Wrong ID"})
			return
		} else {
			c.JSON(http.StatusNotFound, gin.H{"msg": "Failure", "error": err.Error()})
			return
		}
	}
	if err := driver.Delete(&author).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Error deleting author", "error": err.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Successfully deleted author"})
}
