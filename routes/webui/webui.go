package webui

import (
	"log"
	"net/http"

	//web ui
	_ "github.com/germanfilipp/transportation/public/webui"
	"github.com/rakyll/statik/fs"
)

var fsW, fsErr = fs.New()

// GetFiles return handler with fileserver
func GetFiles(path string) http.Handler {
	if fsErr != nil {
		log.Fatal(fsErr)
	}
	return http.StripPrefix(path, http.FileServer(fsW))
}
