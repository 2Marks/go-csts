package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseRequestBody[T any](r *http.Request, payload *T) error {
	body := r.Body

	if body == nil {
		return errors.New("no payload sent in request body")
	}

	return json.NewDecoder(body).Decode(payload)
}

func GetPathRequestIntVal(r *http.Request, key string) int {
	vars := mux.Vars(r)
	value, err := strconv.ParseInt(vars[key], 10, 64)

	if err != nil {
		return 0
	}

	return int(value)
}

func GetQueryRequestIntVal(r *http.Request, key string) int {
	value, err := strconv.ParseInt(r.URL.Query().Get(key), 10, 64)

	if err != nil {
		return 0
	}

	return int(value)
}
