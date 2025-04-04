package controller

import (
	"fmt"
	"io"
	"net/http"
)

func TestConnection(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Couldn't read body", http.StatusBadRequest)
		return
	}

	content := string(body)
	fmt.Println("Received from frontend: ", content)

	// respond to client
	w.Header().Set("Access-Control-ALlow-Origin","*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Received: " + content))
}
