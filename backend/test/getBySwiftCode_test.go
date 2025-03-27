package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

// This test will check if non-existent swift code will get us information, should fail with http.StatusNotFound.
func (s *APITestSuite) TestBySwiftCodeOnNonExistent() {
	req, _ := http.NewRequest("GET", s.path+"/ABCDEFGHIJK", nil)
	w := httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusNotFound, w.Code)
	assert.Contains(s.T(), w.Body.String(), "{\"message\":\"")
}

// This test will check for a headquarters swift code.
func (s *APITestSuite) TestBySwiftCodeHeadquarters() {
	req, _ := http.NewRequest("GET", s.path+"/POLSPLAWXXX", nil)
	w := httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusOK, w.Code)
}

// This test will check for a branch swift code.
func (s *APITestSuite) TestBySwiftCodeBranches() {
	req, _ := http.NewRequest("GET", s.path+"/LITWLTDEADD", nil)
	w := httptest.NewRecorder()
	s.gin.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusOK, w.Code)
}
