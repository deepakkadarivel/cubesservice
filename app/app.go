package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"github.com/sirupsen/logrus"
	"fmt"
)

type App struct {
	Router *mux.Router
	Log    logrus.Logger
}

func (app *App) InitializeRoutes() {
	app.Router.HandleFunc("/", app.homeHandler).Methods("GET")
}
func (app *App) Run(addr string) {
	app.Log.Info("Enabling Routes")
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func (app *App) homeHandler(res http.ResponseWriter, req *http.Request) {
	app.Log.Debug("route '/' triggered")
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, "Welcom to Cubes")
}
