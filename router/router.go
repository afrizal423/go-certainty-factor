package router

import (
	"go_certainty_factor/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/penyakit", controller.ListPenyakit).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/hitung", controller.HitungCF).Methods("POST", "OPTIONS")

	return router
}
