package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		wantKey string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Valid APIKey Authorization header",
			headers: getHeader("ApiKey Valid_key"),
			wantKey: "Valid_key",
			wantErr: false,
		},
		{
			name:    "Missing Authorization header",
			headers: getHeader(""),
			wantKey: "",
			wantErr: true,
			errMsg:  "no authorization header included",
		},
		{
			name:    "Malformed Authorization header",
			headers: getHeader("Invalid_ApiKey"),
			wantKey: "",
			wantErr: true,
			errMsg:  "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotKey != tt.wantKey {
				t.Errorf("GetAPIKey() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}

}

func getHeader(authHeader string) http.Header {
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}

	return req.Header
}
