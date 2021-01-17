package router

import (
	"github.com/gorilla/mux"
	"github.com/sahindagdelen/goserver/middleware"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/graphql", middleware.ExecuteQueryGraphql).Methods("POST", "OPTIONS")
	return router
}
