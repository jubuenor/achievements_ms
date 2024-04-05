package routes

import (
	"github.com/gorilla/mux"

	"github.com/FinanceUN/Achievements/controllers"
)

func IndexRoutes(router *mux.Router) {
	router.HandleFunc("/ping", controllers.PingController).Methods("GET")
}
