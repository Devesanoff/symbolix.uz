package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCyrillicToLatinHandler(t *testing.T) {
	tests := []struct {
		name         string
		reqBody      ConvertRequest
		expectedRes  string
	}{
		{
			name:        "Standard mode explicitly",
			reqBody:     ConvertRequest{Text: "Ўzbekiston, Шahar, Чoʻl", Mode: "standard"},
			expectedRes: "Oʻzbekiston, Shahar, Choʻl",
		},
		{
			name:        "New mode explicitly",
			reqBody:     ConvertRequest{Text: "Ўzbekiston, Шahar, Чoʻl", Mode: "new"},
			expectedRes: "Özbekiston, Şahar, Çoʻl",
		},
		{
			name:        "Empty mode defaults to standard",
			reqBody:     ConvertRequest{Text: "Ўzbekiston", Mode: ""},
			expectedRes: "Oʻzbekiston",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.reqBody)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/convert/krill-lotin", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			CyrillicToLatinHandler(w, req)

			if w.Code != http.StatusOK {
				t.Fatalf("expected status 200, got %d", w.Code)
			}

			var resp ConvertResponse
			if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if resp.Result != tt.expectedRes {
				t.Errorf("expected %q, got %q", tt.expectedRes, resp.Result)
			}
		})
	}
}
