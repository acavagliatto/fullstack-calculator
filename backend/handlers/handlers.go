package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/acavagliatto/fullstack-calculator/backend/calculator"
)

// CalculateRequest represents the incoming calculation request
type CalculateRequest struct {
	Operation string    `json:"operation"`
	Operands  []float64 `json:"operands"`
}

// CalculateResponse represents the calculation response
type CalculateResponse struct {
	Result float64 `json:"result"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// CalculateHandler handles POST /api/calculate requests
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Validate operation
	if req.Operation == "" {
		sendError(w, "operation is required", http.StatusBadRequest)
		return
	}

	// Validate operands
	if req.Operands == nil || len(req.Operands) == 0 {
		sendError(w, "operands are required", http.StatusBadRequest)
		return
	}

	// Perform calculation
	result, err := calculator.Calculate(req.Operation, req.Operands)
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CalculateResponse{Result: result})
}

// HealthHandler handles GET /api/health requests
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func sendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
