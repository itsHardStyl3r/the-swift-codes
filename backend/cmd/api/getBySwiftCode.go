package api

import (
	"github.com/gin-gonic/gin"
	"github.com/itsHardStyl3r/the-swift-codes/internal/models"
	"github.com/itsHardStyl3r/the-swift-codes/internal/tools"
	"net/http"
)

type swiftResponse struct {
	Address       string          `json:"address"`
	BankName      string          `json:"bankName"`
	CountryISO2   string          `json:"countryISO2"`
	CountryName   string          `json:"countryName,omitempty"`
	IsHeadquarter bool            `json:"isHeadquarter"`
	SwiftCode     string          `json:"swiftCode"`
	Branches      []swiftResponse `json:"branches,omitempty"`
}

// Endpoint 1: Retrieve details of a single SWIFT code whether for a headquarters or branches.
// GET: /v1/swift-codes/{swift-code}:
func bySwiftCode(rg *gin.RouterGroup) {
	request := rg.Group("/swift-codes")
	request.GET("/:swift", func(c *gin.Context) {
		swift := c.Param("swift")
		var bic models.Bic
		result := tools.DB.Joins("Bank").Joins("Country").Where("bic = ?", swift).First(&bic)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Swift code not found",
			})
			return
		}
		response := swiftResponse{
			Address:       bic.Address,
			BankName:      bic.Bank.Name,
			CountryISO2:   bic.Country.Iso2,
			CountryName:   bic.Country.Name,
			IsHeadquarter: bic.IsHeadquarter(),
			SwiftCode:     bic.Bic,
		}

		if bic.IsHeadquarter() {
			var bics []models.Bic
			tools.DB.Preload("Bank").Preload("Country").Where("bank_id = ?", bic.BankId).Find(&bics)
			branches := make([]swiftResponse, 0, len(bics))
			for _, branch := range bics {
				branches = append(branches, swiftResponse{
					Address:       branch.Address,
					BankName:      branch.Bank.Name,
					CountryISO2:   branch.Country.Iso2,
					CountryName:   branch.Country.Name,
					IsHeadquarter: branch.IsHeadquarter(),
					SwiftCode:     branch.Bic,
				})
			}
			response.Branches = branches
		}
		c.JSON(http.StatusOK, response)
	})
}
