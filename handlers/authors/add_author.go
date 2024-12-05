package authors

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/ArjunMalhotra07/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddAuthor(c *gin.Context, driver *gorm.DB) {
	var author model.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failure", "error": "Error binding json"})
		return
	}
	authorID, err := exec.Command("uuidgen").Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failure", "error": "Error generating user ID"})
		return
	}
	author.ID = strings.TrimSpace(string(authorID))
	if err := driver.Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failure", "error": "Error creating author"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Success", "error": "Added Author"})
}
