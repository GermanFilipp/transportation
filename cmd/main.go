package main

import (
	"log"
	"net/http"

	"github.com/germanfilipp/transportation/config"
	"github.com/germanfilipp/transportation/routes"
	"github.com/germanfilipp/transportation/utils/logger"
	"github.com/rs/cors"
)

func main() {
	conf := config.GetEnv()

	router := routes.GetRouter()
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(conf.Port, logRequest(handler)))
}

func logRequest(handler http.Handler) http.Handler {
	// logger.NewLogger()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
