package router

import (
	"github.com/friendsofgo/graphiql"
	"github.com/gorilla/mux"
	"github.com/sahindagdelen/goserver/middleware"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/graphql", middleware.ExecuteQueryGraphql).Methods("POST", "OPTIONS")

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/api/graphql")
	if err != nil {
		panic(err)
	}
	router.Handle("/api/graphiql", graphiqlHandler)
	return router
}
