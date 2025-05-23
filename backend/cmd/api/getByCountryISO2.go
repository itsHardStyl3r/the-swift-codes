package api

import (
	"github.com/gin-gonic/gin"
	"github.com/itsHardStyl3r/the-swift-codes/internal/models"
	"github.com/itsHardStyl3r/the-swift-codes/internal/tools"
	"net/http"
)

type countryResponse struct {
	CountryISO2 string          `json:"countryISO2"`
	CountryName string          `json:"countryName"`
	SwiftCodes  []swiftResponse `json:"swiftCodes,omitempty"`
}

// ByCountryCode Endpoint 2: Return all SWIFT codes with details for a specific country (both headquarters and branches).
// GET: /v1/swift-codes/country/{countryISO2code}:
func ByCountryCode(rg *gin.RouterGroup) {
	request := rg.Group("/swift-codes/country")
	request.GET("/:iso2", func(c *gin.Context) {
		swift := c.Param("iso2")
		if len(swift) != 2 { // using len(), since iso2 consists only of a "2-byte rune"
			abortWithJSON(c, http.StatusBadRequest, "Provided ISO2 code is invalid.")
			return
		}
		var country models.Country
		result := tools.DB.First(&country, "iso2 = ?", swift)
		if result.Error != nil {
			abortWithJSON(c, http.StatusNotFound, "Country with such ISO2 code has not been found.")
			return
		}
		var bics []models.Bic
		tools.DB.Joins("Bank").Joins("Country").Where("country_id = ?", country.Id).Find(&bics)
		branches := make([]swiftResponse, 0, len(bics))
		for _, branch := range bics {
			branches = append(branches, swiftResponse{
				Address:       branch.Address,
				BankName:      branch.Bank.Name,
				CountryISO2:   branch.Country.Iso2,
				IsHeadquarter: branch.IsHeadquarter(),
				SwiftCode:     branch.Bic,
			})
		}
		response := countryResponse{
			CountryISO2: country.Iso2,
			CountryName: country.Name,
			SwiftCodes:  branches,
		}
		c.JSON(http.StatusOK, response)
	})
}
