package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()
	expected := `{"message":"pong"}`

	responseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		fmt.Println("no response from /ping")
	}
	router.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, expected, responseRecorder.Body.String())

	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
	// var response map[string]string
	// Err := json.Unmarshal([]byte(responseRecorder.Body.String()), &response)
	// value, exists := response["message"]
	// assert.Nil(t, Err)
	// assert.Equal(t, body["message"], value)
	// assert.True(t, exists)
}