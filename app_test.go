package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	setRoutes()
}

// Test AddNumbers with correct usage
func TestAddNumbers(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/add/2/4", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "6", w.Body.String())
}

// Test AddNumbers where param cannot be converted to int
func TestAddNumbersBadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/add/2/four", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, addNumbersUsageMsg, w.Body.String())
}

// Test SubtractNumbers with correct usage
func TestSubtractNumbers(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/subtract/2/4", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "-2", w.Body.String())
}

// Test SubtractNumbers where param cannot be converted to int
func TestSubtractNumbersBadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/subtract/2/four", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, subtractNumbersUsageMsg, w.Body.String())
}