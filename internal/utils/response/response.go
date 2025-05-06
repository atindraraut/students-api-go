package response

import (
	"encoding/json"
	"net/http"
)

type Response struct{
	Status  string     `json:"status"`
	Error   string     `json:"error"`
}

const (
	StatusOK	 = "OK"
	StatusError	 = "ERROR"
	StatusCreated = "CREATED"
	StatusNotFound = "NOT_FOUND"
	StatusBadRequest = "BAD_REQUEST"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response{
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}
func ValidationError(err error) Response{
	return Response{
		Status: StatusBadRequest,
		Error:  err.Error(),
	}
}