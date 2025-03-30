package test

import (
	"bytes"
	"encoding/json"
	"github.com/itsHardStyl3r/the-swift-codes/cmd/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

// This will work properly as long as PostSwiftRequest json requirements are met, so it doesn't quite matter
// what you put here. Request should fail with http.StatusBadRequest.
func (s *APITestSuite) TestPostSwiftCodeInvalidBody() {
	re := api.PostSwiftRequest{
		Address:       "",
		BankName:      "",
		CountryISO2:   "",
		CountryName:   "",
		IsHeadquarter: false,
		SwiftCode:     "",
	}
	jsonre, _ := json.Marshal(&re)
	req, _ := http.NewRequest("POST", s.path, bytes.NewBuffer(jsonre))
	w := httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

// This should check if method PostSwiftRequest.isValidSwiftCode() fails.
// Naturally, the request would end up with http.StatusBadRequest.
func (s *APITestSuite) TestPostSwiftCodeInvalidSwift() {
	re := api.PostSwiftRequest{
		Address:       "Some address",
		BankName:      "Bank of Poland",
		CountryISO2:   "PL",
		CountryName:   "POLAND",
		IsHeadquarter: false,
		SwiftCode:     "POLSAWAWXX",
	}
	jsonre, _ := json.Marshal(&re)
	req, _ := http.NewRequest("POST", s.path, bytes.NewBuffer(jsonre))
	w := httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

// As stated in the API, this check could get looser later on (like the way new Banks are added), so
// most likely will need refactoring if that happens.
// Right now, the request should fail with http.StatusNotFound.
func (s *APITestSuite) TestPostSwiftCodeInvalidCountry() {
	re := api.PostSwiftRequest{
		Address:       "Some address",
		BankName:      "Bank of Poland",
		CountryISO2:   "PL",
		CountryName:   "POLAN",
		IsHeadquarter: false,
		SwiftCode:     "POLSDDAWXXX",
	}
	jsonre, _ := json.Marshal(&re)
	req, _ := http.NewRequest("POST", s.path, bytes.NewBuffer(jsonre))
	w := httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusNotFound, w.Code)
}

// This contains correct country, but it doesn't match in the swift code.
func (s *APITestSuite) TestPostSwiftCodeCountryMismatch() {
	re := api.PostSwiftRequest{
		Address:       "Some address",
		BankName:      "Bank of Poland",
		CountryISO2:   "PL",
		CountryName:   "POLAND",
		IsHeadquarter: true,
		SwiftCode:     "POLSLPDDXXX",
	}
	jsonre, _ := json.Marshal(&re)
	req, _ := http.NewRequest("POST", s.path, bytes.NewBuffer(jsonre))
	w := httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

// This is supposed to add a branch, but swift code ends with XXX, pointing to headquarter.
func (s *APITestSuite) TestPostSwiftCodeHeadquarterMismatch() {
	re := api.PostSwiftRequest{
		Address:       "Some address",
		BankName:      "Bank of Poland",
		CountryISO2:   "PL",
		CountryName:   "POLAND",
		IsHeadquarter: false,
		SwiftCode:     "POLSLPDDXXX",
	}
	jsonre, _ := json.Marshal(&re)
	req, _ := http.NewRequest("POST", s.path, bytes.NewBuffer(jsonre))
	w := httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

// This contains correct information, so POST should succeed.
func (s *APITestSuite) TestPostSwiftCode() {
	re := api.PostSwiftRequest{
		Address:       "Some address",
		BankName:      "Bank of Poland",
		CountryISO2:   "PL",
		CountryName:   "POLAND",
		IsHeadquarter: true,
		SwiftCode:     "POLSPLDDXXX",
	}
	jsonre, _ := json.Marshal(&re)
	req, _ := http.NewRequest("POST", s.path, bytes.NewBuffer(jsonre))
	w := httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusOK, w.Code)
}
