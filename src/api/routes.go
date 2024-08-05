package api

import (
	"fmt"
	"net/http"

	"github.com/AlexM141200/munros-api/src/model"
)

func handleGetMunros(w http.ResponseWriter, r *http.Request) {
	munro := model.Munro{Id: "1", Name: "Ben Nevis", Height: 1245.0, Location: "Grampian Mountains"}
	res := fmt.Sprintf("Munro Details: ID= %s, Name=%s, Height=%f, Location=%s", munro.Id, munro.Name, munro.Height, munro.Location)
	w.Write([]byte(res))
}

func handleMunroByID(w http.ResponseWriter, r *http.Request) {
	munroID := r.PathValue("id")
	response := fmt.Sprintf("Munro ID: %s", munroID)
	w.Write([]byte(response))
}

func handlePostMunro(w http.ResponseWriter, r *http.Request) {

}
