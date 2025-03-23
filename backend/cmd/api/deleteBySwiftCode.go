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
			abortWithJSON(c, http.StatusBadRequest, "Provided swift code is invalid.")
			return
		}

		result := tools.DB.Where("bic = ?", swift).Delete(&models.Bic{})
		if result.Error != nil {
			abortWithJSON(c, http.StatusInternalServerError, "There's been an error deleting this swift code.")
			return
		}
		if result.RowsAffected == 0 {
			abortWithJSON(c, http.StatusNotFound, "Provided swift code has not been found.")
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Swift code successfully deleted.",
		})
	})
}
