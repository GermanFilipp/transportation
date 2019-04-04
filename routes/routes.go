package routes

import (
	"github.com/germanfilipp/transportation/controllers/transportation"
	"github.com/germanfilipp/transportation/routes/swagger"
	"github.com/germanfilipp/transportation/routes/webui"
	"github.com/gorilla/pat"
)

const (
	transportPath = "/transport-solution"
	swaggerPath   = "/swagger/"
	rootPath      = "/"
)

//GetRouter return routes
func GetRouter() *pat.Router {
	router := pat.New()
	router.Post(transportPath, transportation.Create)
	router.PathPrefix(swaggerPath).Handler(swagger.GetFiles(swaggerPath))
	router.PathPrefix(rootPath).Handler(webui.GetFiles(rootPath))
	return router
}
