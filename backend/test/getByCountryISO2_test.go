package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

// This test will check if we're getting http.StatusBadRequest when ISO2 code size is not two,
// both "A" and "AAA" should fail with http.StatusBadRequest.
func (s *APITestSuite) TestByCountryISO2OnInvalid() {
	req, _ := http.NewRequest("GET", s.path+"/country/A", nil)
	w := httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusBadRequest, w.Code)

	req, _ = http.NewRequest("GET", s.path+"/country/AAA", nil)
	w = httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

// This test will check if correct but non-existent ISO2 code will get us information, should fail
// with http.StatusNotFound.
func (s *APITestSuite) TestByCountryISO2OnNonExistent() {
	req, _ := http.NewRequest("GET", s.path+"/country/EU", nil)
	w := httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusNotFound, w.Code)
}

// This test will check if we can retrieve information about an existing country with their swift codes.
// Not to state the obvious, but please keep in mind that this checks for the actual response body that's inline
// with mock data, any changes made to it will most likely result in this failing.
func (s *APITestSuite) TestByCountryISO2WithSwifts() {
	req, _ := http.NewRequest("GET", s.path+"/country/PL", nil)
	w := httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.Equal(s.T(), w.Body.String(),
		`{"countryISO2":"PL","countryName":"POLAND","swiftCodes":[{"address":"Some address","bankName":"Bank of Poland","countryISO2":"PL","isHeadquarter":true,"swiftCode":"POLSPLAWXXX"}]}`)
}

// This test will check if we can retrieve information about an existing country without swift codes.
func (s *APITestSuite) TestByCountryISO2() {
	req, _ := http.NewRequest("GET", s.path+"/country/BG", nil)
	w := httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.Equal(s.T(), w.Body.String(), `{"countryISO2":"BG","countryName":"BULGARIA"}`)
}
