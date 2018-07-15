package routers

import (
	"github.com/gorilla/mux"
)

//InitRoutes for the app
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the User entity
	router = SetUserRoutes(router)
	// Routes for the Task entity
	router = SetTaskRoutes(router)
	// Routes for the Note entity
	router = SetNoteRoutes(router)

	return router
}
