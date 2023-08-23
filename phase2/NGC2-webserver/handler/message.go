package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type HandleMessage struct {
	Message string
}

func InternalError(w http.ResponseWriter,err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(HandleMessage{Message: "Status Internal Server Error : db"})
	log.Fatal(err)
}