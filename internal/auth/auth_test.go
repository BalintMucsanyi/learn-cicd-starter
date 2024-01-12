package auth // replace with your package name

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

// Your GetAPIKey function here...

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name:          "Valid Authorization Header",
			headers:       http.Header{"Authorization": []string{"ApiKey 12345"}},
			expectedKey:   "12345",
			expectedError: nil,
		},
		{
			name:          "No Authorization Header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name:          "Malformed Authorization Header",
			headers:       http.Header{"Authorization": []string{"InvalidFormat"}},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			key, err := GetAPIKey(test.headers)

			if key != test.expectedKey {
				t.Errorf("Expected key %v, got %v", test.expectedKey, key)
			}

			if !errors.Is(err, test.expectedError) {
				t.Errorf("Expected error %v, got %v", test.expectedError, err)
			}
		})
	}
}
