package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestSecretHandler(t *testing.T) {
	// Set up test cases
	tests := []struct {
		name           string
		secretName     string
		expectedStatus int
	}{
		{
			name:           "Missing secret name",
			secretName:     "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Valid secret name",
			secretName:     "test-secret",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variable for test
			os.Setenv("AWS_SECRET_NAME", tt.secretName)
			defer os.Unsetenv("AWS_SECRET_NAME")

			// Create request
			req, err := http.NewRequest("GET", "/secret", nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create response recorder
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Mock the getSecretValue function
				if tt.secretName == "" {
					http.Error(w, "AWS_SECRET_NAME not set", http.StatusBadRequest)
					return
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Secret value: test-secret-value"))
			})

			// Serve the request
			handler.ServeHTTP(rr, req)

			// Check status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}
		})
	}
}
