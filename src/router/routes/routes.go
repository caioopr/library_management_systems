package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents the API route structure
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// Configure adds every route to the router
func Configure(router *mux.Router) *mux.Router {
	routes := usersRoutes

	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.AuthRequired {
			router.HandleFunc(route.URI, middlewares.Authorize(route.Function)).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI, route.Function).Methods(route.Method)
		}
	}

	return router
}
