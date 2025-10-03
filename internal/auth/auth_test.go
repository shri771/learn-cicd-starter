package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestCheckAPIKey(t *testing.T) {
	// Create Header strut
	tests := []struct {
		name           string
		input          http.Header
		expectedResult string
		expectedErr    error
	}{
		{
			name: "Valid Header",
			input: http.Header{
				"Authorization": []string{"ApiKey d71cc647-ffa9-4252-a2ea-786ac273ce46"},
			},
			expectedResult: "d71cc647-ffa9-4252-a2ea-786ac273ce46",
			expectedErr:    nil,
		},
		{
			name: "Missing API Header",
			input: http.Header{
				"Authorization": []string{""},
			},
			expectedResult: "",
			expectedErr:    ErrNoAuthHeaderIncluded,
		},
		{
			name: "Missing API Key",
			input: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedResult: "",
			expectedErr:    errors.New("malformed authorization header"),
		},
	}
	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.input)

			if apiKey != tt.expectedResult {
				t.Errorf("GetAPIKey() input: %v, expected: %v", tt.input, tt.expectedResult)
				return
			}

			if (tt.expectedErr != nil) && err == nil {
				t.Errorf("GetAPIKey() expected error: %v, but got nil", tt.expectedErr)
			}

			if tt.expectedErr == nil && err != nil {
				t.Errorf("GetAPIKey() expected no error, but got: %v", err)
			}
		})
	}
}
