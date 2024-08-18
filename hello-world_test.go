package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloServer(t *testing.T) {
	tests := []struct {
		name           string
		requestPath    string
		expectedStatus int
		expectedBody   string
	}{
		{"root path", "/", http.StatusOK, "Hello, !"},
		{"name path", "/Osher", http.StatusOK, "Hello, Osher!"},
		{"empty name path", "//", http.StatusOK, "Hello, !"},
		{"trailing slash", "/Osher/", http.StatusOK, "Hello, Osher/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request to pass to our handler
			req := httptest.NewRequest("GET", tt.requestPath, nil)
			
			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Call the HelloServer function with the request and response recorder
			HelloServer(rr, req)

			// Check the status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			// Check the response body
			if rr.Body.String() != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expectedBody)
			}
		})
	}
}
