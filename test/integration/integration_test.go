// Integration Tests for Dating App APIs in Go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const baseUrl = "http://localhost:8080/api/v1"

var authToken string

func TestIntegration(t *testing.T) {
	t.Run("Sign Up", func(t *testing.T) {
		requestBody := map[string]interface{}{
			"email":             "testuser@gmail.com",
			"name":              "Test User",
			"date_of_birth":     "1990-01-01",
			"gender":            "male",
			"profile_image_url": "https://example.com/profile.jpg",
			"description":       "Integration Test User",
			"password":          "password123",
		}
		body, _ := json.Marshal(requestBody)
		response, err := http.Post(baseUrl+"/auth/sign-up", "application/json", bytes.NewBuffer(body))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)

		var responseBody map[string]interface{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		assert.Contains(t, responseBody, "user_id")
	})

	t.Run("Login", func(t *testing.T) {
		requestBody := map[string]interface{}{
			"email":    "testuser@gmail.com",
			"password": "password123",
		}
		body, _ := json.Marshal(requestBody)
		response, err := http.Post(baseUrl+"/auth/login", "application/json", bytes.NewBuffer(body))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)

		var responseBody map[string]interface{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		assert.Contains(t, responseBody, "token")
		authToken = responseBody["token"].(string)
	})

	t.Run("Get Recommendations", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseUrl+"/recommendations", nil)
		req.Header.Set("Authorization", "Bearer "+authToken)
		response, err := http.DefaultClient.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)

		var responseBody []interface{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		assert.IsType(t, []interface{}{}, responseBody)
	})

	t.Run("Get Profile", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseUrl+"/auth/profile", nil)
		req.Header.Set("Authorization", "Bearer "+authToken)
		response, err := http.DefaultClient.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)

		var responseBody map[string]interface{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		assert.Equal(t, "testuser@gmail.com", responseBody["email"])
		assert.Equal(t, "Test User", responseBody["name"])
	})

	t.Run("Swipe", func(t *testing.T) {
		requestBody := map[string]interface{}{
			"swiped_user_id": "example-user-id",
			"swipe_type":     "pass",
		}
		body, _ := json.Marshal(requestBody)
		req, _ := http.NewRequest("POST", baseUrl+"/swipe", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer "+authToken)
		req.Header.Set("Content-Type", "application/json")
		response, err := http.DefaultClient.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)

		var responseBody map[string]interface{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		assert.Equal(t, "success", responseBody["status"])
	})

	t.Run("Upgrade Premium", func(t *testing.T) {
		req, _ := http.NewRequest("POST", baseUrl+"/premium", nil)
		req.Header.Set("Authorization", "Bearer "+authToken)
		response, err := http.DefaultClient.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)

		var responseBody map[string]interface{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		assert.Equal(t, "active", responseBody["premium_status"])
	})
}
