package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  string
		error bool
	}{
		"simple": {
			input: http.Header{
				"Authorization": []string{"ApiKey authToken"},
			},
			want:  "authToken",
			error: false,
		},
		"no auth header": {
			input: http.Header{},
			want:  "",
			error: true,
		},
		"malformed auth header 1": {
			input: http.Header{
				"Authorization": []string{"authToken"},
			},
			want: "",
			error: true,
		},
		"malformed auth header 2": {
			input: http.Header{
				"Authorization": []string{"Bearer token"},
			},
			want: "",
			error: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if tc.error && err == nil {
				t.Fatalf("Expected error")
			}
			if !tc.error && err != nil {
				t.Fatalf("Did not expect error")
			}
			if tc.want != got {
				t.Fatalf(fmt.Sprintf("Expected %s, got %s", tc.want, got))
			}
		})
	}
}
