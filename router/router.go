package router

import (
	"go_certainty_factor/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/penyakit", controller.ListPenyakit).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/hitung", controller.HitungCF).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/buku/{id}", controller.AmbilBuku).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/buku", controller.TmbhBuku).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/buku/{id}", controller.UpdateBuku).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/api/buku/{id}", controller.HapusBuku).Methods("DELETE", "OPTIONS")

	return router
}
