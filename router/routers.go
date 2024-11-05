package router

import (
	"encoding/json"
	"net/http"

	"github.com/anujmritunjay/go-postgres/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// router.HandleFunc("/api/stock/{id}", ).Methods("GET", "OPTIONS")
	router.HandleFunc("/create-stock", middleware.CreateStock).Methods("POST")
	router.HandleFunc("/", rootRoute).Methods("GET")
	// router.HandleFunc("/")
	// router.HandleFunc("/")
	return router

}

func rootRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "message": "Hello from the Go lang server"})
}
