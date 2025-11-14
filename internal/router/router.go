package router

import (
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/Etevaldo15/go-nif-validator/docs" // swag docs
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/Etevaldo15/go-nif-validator/internal/api/handlers"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}).Methods("GET")

	// API v1
	r.HandleFunc("/api/v1/validate-nif/{nif}", handlers.ValidateNIF).Methods("GET")

	// Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}

