package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParams struct {
	Username string `json:"username"`
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Success Code, Usually 200
	Code int

	// Account Balance
	Balance uint64
}

// Error Response
type Error struct {
	// Error code
	Code int

	// Error Message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// Write JSON Response
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Internal Server Error", http.StatusInternalServerError)
	}
)
