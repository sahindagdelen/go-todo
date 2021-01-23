package main

import (
	"fmt"
	"github.com/sahindagdelen/go-todo/api/server"
	"log"
	"net/http"
)

func main() {
	r := server.Router()
	fmt.Println("Starting server on the port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
