package main

import (
	"log"
	"net/http"
	"os"

	"github.com/germanfilipp/transportation/routes"
	"github.com/germanfilipp/transportation/utils/logger"
	"github.com/rs/cors"
)

func main() {
	// conf := config.GetEnv()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := routes.GetRouter()
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, logRequest(handler)))
}

func logRequest(handler http.Handler) http.Handler {
	// logger.NewLogger()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
