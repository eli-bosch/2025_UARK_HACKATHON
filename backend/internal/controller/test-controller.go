package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func TestConnection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contentTitle := vars["content"]

	fmt.Println(contentTitle)

}
