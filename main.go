package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anujmritunjay/go-postgres/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting the server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
