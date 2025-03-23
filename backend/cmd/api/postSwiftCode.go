package api

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/itsHardStyl3r/the-swift-codes/internal/models"
	"github.com/itsHardStyl3r/the-swift-codes/internal/tools"
	"net/http"
)

type postSwiftRequest struct {
	Address       string `json:"address" binding:"required"`
	BankName      string `json:"bankName" binding:"required"`
	CountryISO2   string `json:"countryISO2" binding:"required"`
	CountryName   string `json:"countryName" binding:"required"`
	IsHeadquarter bool   `json:"isHeadquarter" binding:"required"`
	SwiftCode     string `json:"swiftCode" binding:"required"`
}

func (PostSwiftRequest postSwiftRequest) getBankCode() string {
	return PostSwiftRequest.SwiftCode[0:4]
}

// Endpoint 3: Adds new SWIFT code entries to the database for a specific country.
// POST: /v1/swift-codes
func postSwiftCode(rg *gin.RouterGroup) {
	request := rg.Group("/swift-codes")
	request.POST("", func(c *gin.Context) {
		var body postSwiftRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		var bic models.Bic
		if tools.DB.Where("bic = ?", body.SwiftCode).First(&bic).RowsAffected > 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "This swift code is already in use."})
			return
		}

		var country models.Country
		if err := tools.DB.Where("iso2 = ? AND name = ?",
			body.CountryISO2, body.CountryName).First(&country).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Country with such iso2 has not been found or there is a mismatch between ISO2 and country name.",
			})
			return
		}

		var bank models.Bank
		if err := tools.DB.Where("bank_code = ?", body.getBankCode()).First(&bank).Error; err != nil {
			log.Warn("This bank is not in the database yet. Adding...")
			bank = models.Bank{
				Name:     body.BankName,
				BankCode: body.getBankCode(),
			}
			if err := tools.DB.Create(&bank).Error; err != nil {
				log.Errorf("Failed to create bank: %v", err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "There was an error adding new bank: " + err.Error(),
				})
				return
			}
		}

		newBic := models.Bic{
			Bic:          body.SwiftCode,
			BankId:       bank.Id,
			CountryId:    country.Id,
			LocationCode: body.SwiftCode[6:8],
			Branch:       body.SwiftCode[8:11],
			Address:      body.Address,
		}

		if err := tools.DB.Create(&newBic).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to push new BIC to the database.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "New BIC code added."})
	})
}
