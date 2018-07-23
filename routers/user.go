package routers

import (
	"github.com/gorilla/mux"
	"github.com/vsfastrack/TaskManager/controllers"
)

//SetUserRoutes handlers for user
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	return router
}
