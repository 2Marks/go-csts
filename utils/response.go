package utils

import (
	"encoding/json"
	"net/http"
)

func WriteResponseToJson[T any](w http.ResponseWriter, statusCode int, v T) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(v)
}

func WriteErrorToJson(w http.ResponseWriter, statusCode int, err error) error {
	return WriteResponseToJson(
		w,
		statusCode,
		map[string]any{"success": false, "message": err.Error()},
	)
}

func WriteCreateSuccessResponse[T any](w http.ResponseWriter, v T) error {
	jsonResponse := map[string]any{
		"success": true,
		"message": "resource created successfully",
		"data":    v,
	}

	return WriteResponseToJson(w, http.StatusCreated, jsonResponse)
}

func WriteCreateOkResponse[T any](w http.ResponseWriter, v T) error {
	jsonResponse := map[string]any{
		"success": true,
		"message": "resource fetched successfully",
		"data":    v,
	}

	return WriteResponseToJson(w, http.StatusOK, jsonResponse)
}
