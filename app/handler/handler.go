package handler

import (
	"encoding/json"
	"errors"
	"net/http"
)

func unmarshallRequest(requestStruct interface{}, w http.ResponseWriter, r *http.Request) error {
	err := errors.New("missing request body")
	if r.Body != nil {
		err = json.NewDecoder(r.Body).Decode(&requestStruct)
	}
	if err == nil {
		return nil
	}
	http.Error(w, "bad request", http.StatusBadRequest)
	return err
}

type errorMessage struct {
	Message string `json:"message"`
}

func marshallErrorResponse(message string, w http.ResponseWriter) {
	ret, err := json.Marshal(errorMessage{
		Message: message,
	})

	if err != nil {
		http.Error(w, `{"message": "Error on writeErrorMessage"}`, http.StatusBadRequest)
		return
	}

	http.Error(w, string(ret), http.StatusBadRequest)
}

func marshallResponse(res interface{}, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(&res)
}
