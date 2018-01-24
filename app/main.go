package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
)

type App struct {
	Router *mux.Router
}

func (app *App) initializeRoutes() {
	app.Router.Handle("/", homeHandler).Methods("GET")
}
func (app *App) Run(addr string) {
	fmt.Println("Enabeling Routes")
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func main() {
	a := App{mux.NewRouter()}
	a.initializeRoutes()
	a.Run(":7700")
}

var homeHandler = http.HandlerFunc(
	func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		fmt.Fprintln(res, "Welcom to Cubes")
	},
)
