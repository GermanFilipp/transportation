package logger

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var (
	//Log instance var
	Log      = logrus.New()
	fileName = envName() + ".log"
	logpath  = "/tmp/"
)

func init() {
	Log.SetFormatter(&log.JSONFormatter{})
}

//Info log in stdout and file
func Info(v ...interface{}) {
	file := getFile()
	defer file.Close()
	logrus.Info(v...) // log to os.Stdout
	Log.Out = file
	Log.Info(v...) // log to file
}

//Warn log in stdout and file
func Warn(v ...interface{}) {
	file := getFile()
	defer file.Close()
	logrus.Warn(v...)
	Log.Out = file
	Log.Warn(v...)
}

//Error log in stdout and file
func Error(v ...interface{}) {
	file := getFile()
	defer file.Close()
	logrus.Error(v...)
	Log.Out = file
	Log.Error(v...)
}

//Fatal log in stdout and file
func Fatal(v ...interface{}) {
	file := getFile()
	defer file.Close()
	logrus.Fatal(v...)
	Log.Out = file
	Log.Fatal(v...)
}

func envName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	return env
}

func getFile() *os.File {
	p, _ := os.Getwd()
	currentPath := p + logpath
	if _, err := os.Stat(currentPath); os.IsNotExist(err) {
		os.Mkdir(currentPath, os.ModePerm)
	}

	file, err := os.OpenFile(filepath.Join(currentPath, filepath.Base(fileName)), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return file
}
