package main

import (
	"fmt"
	"log"
	"net/http"
	"myjsondb/data_in/handler"
)

func main() {
	// Start the Data In HTTP service
	http.HandleFunc("/data_in", handler.DataInHandler)
	fmt.Println("Data In service started on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

