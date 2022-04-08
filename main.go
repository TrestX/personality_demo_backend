package main

import (
	"log"
	"net/http"
	"personality_demo_backend/pkg/router"

	"github.com/aekam27/trestCommon"
	"github.com/rs/cors"
)

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.AllowAll().Handler
	return handleCORS(handler)
}

func main() {

	trestCommon.LoadConfig()
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":6020", setupGlobalMiddleware(router)))
}
