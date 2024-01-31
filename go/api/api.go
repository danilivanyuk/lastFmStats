package api

import (
	"encoding/json"
	"net/http"
)


type TestStatRequestParams struct {
	Username string
}

type TestStatReponseParams struct {
	// Code of response, 200, 400, etc...
	Code int
	AmountOfTracks int64
}

type Error struct {
	Code int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected Error occured.", http.StatusInternalServerError)
	}
)

