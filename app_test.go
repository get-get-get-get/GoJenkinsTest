package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test AddNumbers with correct usage
func TestAddNumbers(t *testing.T) {
	setRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/add/2/4", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "6", w.Body.String())
}

// Test AddNumbers where param cannot be converted to int
func TestAddNumbersBadRequest(t *testing.T) {
	setRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/add/2/four", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, addNumbersUsageMsg, w.Body.String())

}