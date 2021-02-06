package main

import (
	"fmt"
	"github.com/sahindagdelen/go-todo/api/server"
	"github.com/sahindagdelen/go-todo/config"
	"log"
	"net/http"
)

func main() {
	config.Initialize()
	r := server.Router()
	fmt.Println("Starting server on the port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
