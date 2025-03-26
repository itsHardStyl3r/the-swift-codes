package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

// This test will fail with http.StatusBadRequest, because the code is simply not a swift code.
func (s *APITestSuite) TestDeleteSwiftOnInvalid() {
	req, _ := http.NewRequest("DELETE", s.path+"/1234", nil)
	w := httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

// This test will fail with http.StatusNotFound, because the code is theoretically correct
// but does not exist.
func (s *APITestSuite) TestDeleteSwiftOnNonExistent() {
	req, _ := http.NewRequest("DELETE", s.path+"/XXXXXXXXXXX", nil)
	w := httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusNotFound, w.Code)
}

// This test will delete an existing Swift Code and then verify the {"message":"..."} response structure
// as stated in the requirements.
func (s *APITestSuite) TestDeleteSwiftDeleteBySwiftCode() {
	req, _ := http.NewRequest("DELETE", s.path+"/LITWLTDEXXX", nil)
	w := httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.Contains(s.T(), w.Body.String(), "{\"message\":\"")
}
