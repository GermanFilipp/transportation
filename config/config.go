package config

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tkanos/gonfig"
)

//Configuration store configuration
type Configuration struct {
	Port string `env:"APP_PORT"`
	Env  string `env:"EnvName"`
}

//GetEnv return Configuration with env variables
func GetEnv() Configuration {
	configuration := Configuration{}
	if os.Getenv("ENV") != "production" {
		err := gonfig.GetConf(getFileName(), &configuration)
		if err != nil {
			panic(err)
		}
	}
	port := os.Getenv("PORT")
	if port != "" {
		configuration.Port = ":" + port
	}
	return configuration
}

func getFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	filename := []string{"config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))
	return filePath
}
