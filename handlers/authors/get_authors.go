package authors

import (
	"net/http"

	"github.com/ArjunMalhotra07/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAuthors(c *gin.Context, driver *gorm.DB) {
	var authors []model.Author
	if err := driver.Preload("Books").Find(&authors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failure", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Success", "authors": authors})
}
