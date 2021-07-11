package main

import (
	"fmt"
	"log"
	"net/http"

	"go_api/router"
)

func main() {
	fmt.Println("Rest API v2.0 - GO ORM")
	r := router.Router()

	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
