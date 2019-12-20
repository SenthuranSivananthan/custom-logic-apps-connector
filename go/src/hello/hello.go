package main

import (
	"encoding/json"
	"net/http"
)

type EchoMessage struct {
	Message string
}

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/api/echo", Echo)
	http.ListenAndServe(":8080", handler)
}

func Echo(w http.ResponseWriter, r *http.Request) {
	var echo EchoMessage

	err := json.NewDecoder(r.Body).Decode(&echo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var echoReply EchoMessage
	echoReply.Message = "Echo: " + echo.Message
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(echoReply)
}