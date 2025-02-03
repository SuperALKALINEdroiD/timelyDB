package handlers

import (
	"fmt"
	"net/http"
)

func InsertHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Insert Endpoint WIP")
}
