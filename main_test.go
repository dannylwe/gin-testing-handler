package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	var response map[string]string
	expected := map[string]string{"message":"pong"}

	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		fmt.Println("no response from /ping")
	}
	router.ServeHTTP(responseRecorder, req)

	Err := json.Unmarshal([]byte(responseRecorder.Body.String()), &response)
	value, exists := response["message"]

	assert.Nil(t, Err)
	assert.True(t, exists)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, expected["message"], value)
	
}