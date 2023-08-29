package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct{
	ResponseJSON interface{} `json:"response_json"`
}

func WriteResponse(w http.ResponseWriter,status int ,data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(Response{ResponseJSON: data})
	if err != nil {
		InternalErrorResponse(w,err)
	}
}

func InternalErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(Response{ResponseJSON: "Internal server error, please try again"})
	// check err from cli
	fmt.Println(err)
}
