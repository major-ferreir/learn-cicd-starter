package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		header      string
		expectedKey string
		expectError bool
	}{
		{
			name:        "valid api key",
			header:      "ApiKey abc123",
			expectedKey: "abc123",
			expectError: false,
		},
		{
			name:        "missing authorization header",
			header:      "",
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "wrong authorization type",
			header:      "Bearer abc123",
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "malformed header",
			header:      "ApiKey",
			expectedKey: "",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			headers := http.Header{}

			if tc.header != "" {
				headers.Set("Authorization", tc.header)
			}

			key, err := GetAPIKey(headers)

			if tc.expectError {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if key != tc.expectedKey {
				t.Errorf("expected %q, got %q", tc.expectedKey, key)
			}
		})
	}
}
