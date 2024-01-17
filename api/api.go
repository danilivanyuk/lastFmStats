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
