package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents the API route structure
type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	AuthRequest bool
}

// Configure adds every routes ton the router
func Configure(router *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return router
}
