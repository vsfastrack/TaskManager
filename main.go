package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/vsfastrack/TaskManager/common"
	"github.com/vsfastrack/TaskManager/routers"
)

func main() {
	//Calls startup logic
	common.StartUp()

	//Get the mux router object
	router := routers.InitRoutes()

	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()

}
