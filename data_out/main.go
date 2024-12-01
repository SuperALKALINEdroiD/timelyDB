package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Start the Data In HTTP service
	fmt.Println("Data In service started on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

