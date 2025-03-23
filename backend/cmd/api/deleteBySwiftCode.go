package api

import (
	"github.com/gin-gonic/gin"
	"github.com/itsHardStyl3r/the-swift-codes/internal/models"
	"github.com/itsHardStyl3r/the-swift-codes/internal/tools"
	"net/http"
)

// Endpoint 4: Deletes swift-code data if swiftCode matches the one in the database.
// DELETE: /v1/swift-codes/{swift-code}:
func deleteBySwiftCode(rg *gin.RouterGroup) {
	request := rg.Group("/swift-codes")
	request.DELETE("/:swift", func(c *gin.Context) {
		swift := c.Param("swift")
		if len(swift) != 11 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid swift code."})
			return
		}

		result := tools.DB.Where("bic = ?", swift).Delete(&models.Bic{})
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Error while deleting swift code.",
			})
			return
		}
		if result.RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Swift code has not been found.",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Swift code deleted.",
		})
	})
}
