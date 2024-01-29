package main

import (
	"Book_api/pkg/routers"
	"log"
	"net/http"
)

func main() {

	router := routers.InitializeRoute()

	log.Fatal(http.ListenAndServe(":8181", router))
}
