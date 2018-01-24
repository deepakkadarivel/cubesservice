package main

import (
	"github.com/gorilla/mux"
	"cubesservice/app"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	a := app.App{mux.NewRouter(), log.Logger{Formatter: &log.JSONFormatter{}, Out: os.Stdout, Level: log.DebugLevel}}
	a.InitializeRoutes()
	a.Run(":7500")
}
