package api

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/itsHardStyl3r/the-swift-codes/internal/models"
	"github.com/itsHardStyl3r/the-swift-codes/internal/tools"
	"net/http"
	"strings"
)

type PostSwiftRequest struct {
	Address       string `json:"address" binding:"required"`
	BankName      string `json:"bankName" binding:"required"`
	CountryISO2   string `json:"countryISO2" binding:"required"`
	CountryName   string `json:"countryName" binding:"required"`
	IsHeadquarter bool   `json:"isHeadquarter"`
	SwiftCode     string `json:"swiftCode" binding:"required"`
}

func (PostSwiftRequest PostSwiftRequest) getBankCode() string {
	return PostSwiftRequest.SwiftCode[0:4]
}

// This fucking checks rune length since swift codes don't use special chars.
// Note that this will work only if you use XXX for headquarters, so use BIC11s.
func (PostSwiftRequest PostSwiftRequest) isValidSwiftCode() bool {
	return len(PostSwiftRequest.SwiftCode) == 11
}

// PostSwiftCode Endpoint 3: Adds new SWIFT code entries to the database for a specific country.
// POST: /v1/swift-codes
func PostSwiftCode(rg *gin.RouterGroup) {
	request := rg.Group("/swift-codes")
	request.POST("", func(c *gin.Context) {
		var body PostSwiftRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			log.Debugf("Invalid JSON body: %s.", err)
			abortWithJSON(c, http.StatusBadRequest,
				"Provided JSON request body is invalid. Please check specification and try again.")
			return
		}

		if !body.isValidSwiftCode() {
			abortWithJSON(c, http.StatusBadRequest, "Provided swift code is invalid.")
			return
		}

		var bic models.Bic
		if tools.DB.Where("bic = ?", body.SwiftCode).First(&bic).RowsAffected > 0 {
			abortWithJSON(c, http.StatusBadRequest, "This swift code is already in use.")
			return
		}

		// Note: This does not have to be so strict.
		var country models.Country
		if err := tools.DB.Where("iso2 = ? AND name = ?",
			body.CountryISO2, body.CountryName).First(&country).Error; err != nil {
			abortWithJSON(c, http.StatusNotFound,
				"Country with such iso2 code has not been found or there is a mismatch between it and its country name.")
			return
		}

		var bank models.Bank
		if err := tools.DB.Where("bank_code = ?", body.getBankCode()).First(&bank).Error; err != nil {
			log.Info("This bank is not in the database yet. Adding...")
			bank = models.Bank{
				Name:     body.BankName,
				BankCode: body.getBankCode(),
			}
			if err := tools.DB.Create(&bank).Error; err != nil {
				log.Errorf("Failed to create new bank: %v.", err.Error())
				abortWithJSON(c, http.StatusInternalServerError, "There was an error adding new bank.")
				return
			}
		}

		if !strings.HasPrefix(body.SwiftCode, bank.BankCode+country.Iso2) {
			abortWithJSON(c, http.StatusBadRequest, "Swift code doesn't match the provided country information.")
			return
		}
		if body.IsHeadquarter && !strings.HasSuffix(body.SwiftCode, "XXX") {
			abortWithJSON(c, http.StatusBadRequest, "Swift code is supposed to be headquarters, but is not.")
			return
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
			log.Errorf("Failed to create new bic: %v.", err.Error())
			abortWithJSON(c, http.StatusInternalServerError, "Failed to save new BIC data.")
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "New BIC code added."})
	})
}
