package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"runtime"
	"runtime/debug"
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

	errMsg := errorMessage{
		Message: message,
		Func:    fn,
		Line:    line,
	}

	resAsBytes, err := json.Marshal(errMsg)
	res := string(resAsBytes)

	debug.PrintStack()

	if err != nil {
		http.Error(w, `{"message": "Error on writeErrorMessage"}`, http.StatusBadRequest)
		return
	}

	log.Print("[ERROR]", res)

	http.Error(w, res, http.StatusBadRequest)
}

func marshallResponse(res interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}
