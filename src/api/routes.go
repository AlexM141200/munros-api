package api

import (
	"fmt"
	"net/http"
	//	"github.com/AlexM141200/munros-api/src/model"
)

func handleGetMunros(w http.ResponseWriter, r *http.Request) {

}

func handleMunroByID(w http.ResponseWriter, r *http.Request) {
	munroID := r.PathValue("id")
	response := fmt.Sprintf("Munro ID: %s", munroID)
	w.Write([]byte(response))
}

func () handleMunrosCSV(w http.ResponseWriter, r *http.Request) {

}

func handleGetAllMunros( ) {

}

// ###########################################
// Handling Pages
// ###########################################
func handleIndex(w http.ResponseWriter, r *http.Request) {
	//Server the index.html file
	http.ServeFile(w, r, "./frontend/index.html")
}
