package swagger

import (
	"log"
	"net/http"

	//swagger docs
	_ "github.com/germanfilipp/transportation/public/swagger"
	"github.com/rakyll/statik/fs"
)

var fsFile, fsErr = fs.New()

// GetFiles return handler with fileserver
func GetFiles(path string) http.Handler {
	if fsErr != nil {
		log.Fatal(fsErr)
	}
	return http.StripPrefix(path, http.FileServer(fsFile))
}
