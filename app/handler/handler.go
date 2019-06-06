package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"runtime"
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
	Line    int    `json:"line"`
	Func    string `json:"func"`
}

func marshallErrorResponse(message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	_, fn, line, _ := runtime.Caller(1)

	ret, err := json.Marshal(errorMessage{
		Message: message,
		Func:    fn,
		Line:    line,
	})

	if err != nil {
		http.Error(w, `{"message": "Error on writeErrorMessage"}`, http.StatusBadRequest)
		return
	}

	http.Error(w, string(ret), http.StatusBadRequest)
}

func marshallResponse(res interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}
