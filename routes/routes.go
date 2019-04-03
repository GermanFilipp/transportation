package routes

import (
	"log"
	"net/http"

	"github.com/germanfilipp/transportation/controllers/transportation"
	//swaggerui
	_ "github.com/germanfilipp/transportation/statik"
	"github.com/gorilla/pat"
	"github.com/rakyll/statik/fs"
)

const (
	transportPath = "/transport-solution"
	swaggerPath   = "/swagger/"
)

//GetRouter return routes
func GetRouter() *pat.Router {
	router := pat.New()
	router.Post(transportPath, transportation.Create)
	router.PathPrefix(swaggerPath).Handler(getStaticFiles())
	return router
}

func getStaticFiles() http.Handler {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	return http.StripPrefix(swaggerPath, http.FileServer(statikFS))
}

