package main

import (
	"net/http"
)

func main() {
	initEnvironment()

	router := initRouter()
	http.ListenAndServe(":7001", router)
}

