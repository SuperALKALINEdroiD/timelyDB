package handlers

import (
	"fmt"
	"net/http"

	"github.com/SuperALKALINEdroiD/timelyDB/utils/logs"
)

func InsertHandler(w http.ResponseWriter, r *http.Request) {

	logs.AddWalEntry() // TODO: send data to be logged into WALfor later reconctruction
	// write data to store

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Insert Endpoint WIP")
}
