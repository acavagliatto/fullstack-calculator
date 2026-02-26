package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           interface{}
		expectedStatus int
		expectedResult *float64
		expectedError  bool
	}{
		{
			name:           "valid addition",
			method:         http.MethodPost,
			body:           CalculateRequest{Operation: "add", Operands: []float64{5, 3}},
			expectedStatus: http.StatusOK,
			expectedResult: ptr(8.0),
			expectedError:  false,
		},
		{
			name:           "valid division",
			method:         http.MethodPost,
			body:           CalculateRequest{Operation: "divide", Operands: []float64{10, 2}},
			expectedStatus: http.StatusOK,
			expectedResult: ptr(5.0),
			expectedError:  false,
		},
		{
			name:           "division by zero",
			method:         http.MethodPost,
			body:           CalculateRequest{Operation: "divide", Operands: []float64{10, 0}},
			expectedStatus: http.StatusBadRequest,
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:           "invalid operation",
			method:         http.MethodPost,
			body:           CalculateRequest{Operation: "invalid", Operands: []float64{5, 3}},
			expectedStatus: http.StatusBadRequest,
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:           "missing operation",
			method:         http.MethodPost,
			body:           CalculateRequest{Operands: []float64{5, 3}},
			expectedStatus: http.StatusBadRequest,
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:           "missing operands",
			method:         http.MethodPost,
			body:           CalculateRequest{Operation: "add"},
			expectedStatus: http.StatusBadRequest,
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:           "empty operands",
			method:         http.MethodPost,
			body:           CalculateRequest{Operation: "add", Operands: []float64{}},
			expectedStatus: http.StatusBadRequest,
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:           "invalid JSON",
			method:         http.MethodPost,
			body:           "invalid json",
			expectedStatus: http.StatusBadRequest,
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:           "wrong HTTP method",
			method:         http.MethodGet,
			body:           nil,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:           "square root of negative",
			method:         http.MethodPost,
			body:           CalculateRequest{Operation: "sqrt", Operands: []float64{-4}},
			expectedStatus: http.StatusBadRequest,
			expectedResult: nil,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			var err error

			if tt.body != nil {
				if str, ok := tt.body.(string); ok {
					body = []byte(str)
				} else {
					body, err = json.Marshal(tt.body)
					if err != nil {
						t.Fatalf("failed to marshal request body: %v", err)
					}
				}
			}

			req := httptest.NewRequest(tt.method, "/api/calculate", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			CalculateHandler(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedError {
				var errResp ErrorResponse
				if err := json.NewDecoder(w.Body).Decode(&errResp); err != nil {
					t.Errorf("failed to decode error response: %v", err)
				}
				if errResp.Error == "" {
					t.Error("expected error message in response")
				}
			} else if tt.expectedResult != nil {
				var resp CalculateResponse
				if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
					t.Errorf("failed to decode response: %v", err)
				}
				if resp.Result != *tt.expectedResult {
					t.Errorf("expected result %v, got %v", *tt.expectedResult, resp.Result)
				}
			}
		})
	}
}

func TestHealthHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
	}{
		{
			name:           "valid health check",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "wrong method",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/api/health", nil)
			w := httptest.NewRecorder()

			HealthHandler(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				var resp map[string]string
				if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
					t.Errorf("failed to decode response: %v", err)
				}
				if resp["status"] != "ok" {
					t.Errorf("expected status ok, got %s", resp["status"])
				}
			}
		})
	}
}

func ptr(f float64) *float64 {
	return &f
}
