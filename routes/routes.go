package routes

import (
	"github.com/afikrim/music-app/ctr"
	"github.com/gorilla/mux"
)

type Route struct {
	HomeCTR ctr.Home
}

type RouteInterface interface {
	Routes() *mux.Router
}

func (r *Route) Routes() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", r.HomeCTR.Home).Methods("GET")

	return route
}
