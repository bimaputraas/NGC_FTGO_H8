package handler

import (
	"encoding/json"
	"net/http"
)

type HandleMessage struct {
	Message string
}

func InternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(HandleMessage{Message: "Status Internal Server Error"})
}