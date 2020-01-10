package main

import (
	"log"
	"net/http"

	"github.com/ashmilhussain/goRestAPI_Sample/routers"
)

func main() {
	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
