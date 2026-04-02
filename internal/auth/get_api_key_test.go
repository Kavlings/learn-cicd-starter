package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name: "valid auth header",
			headers: http.Header{
				"Authorization": []string{"ApiKey mykey123"},
			},
			want:    "mykey123",
			wantErr: false,
		},
		{
			name:    "no auth header",
			headers: http.Header{},
			wantErr: true,
		},
		{
			name: "malformed auth header - no ApiKey prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer token123"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
