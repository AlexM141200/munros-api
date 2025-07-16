package handlers

import (
	"net/http"

	"github.com/AlexM141200/munros-api/src/routes"
)

func SetupMunroRoutes(router *http.ServeMux) {

	router.HandleFunc("/api/munros", routes.HandleGetMunros)
	router.HandleFunc("/api/munros/{id}", routes.HandleMunroByID)
	router.HandleFunc("/api/munros/csv", routes.HandleMunrosCSV)
	router.HandleFunc("/api/munros/all", routes.HandleGetAllMunros)
}

func SetupFrontendRoutes(router *http.ServeMux) {
	router.HandleFunc("/", routes.HandleIndex)
}
