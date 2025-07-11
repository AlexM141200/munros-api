package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/AlexM141200/munros-api/src/csv"
	"github.com/AlexM141200/munros-api/src/model"
)

// DataService interface for flexible data source (CSV now, database later)
type DataService interface {
	ReadMunros() ([]model.Munro, error)
}

// Global data service - can be switched between CSV and database
var dataService DataService

func init() {
	// Initialize with CSV service - this can be changed to database later
	dataService = csv.NewCSVService("./data/munrotab_v8.0.1.csv")
}

// CORS middleware
func enableCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
}

// Handle JSON response with error handling
func writeJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// Get all munros with optional filtering
func handleGetMunros(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	munros, err := dataService.ReadMunros()
	if err != nil {
		log.Printf("Error reading munros: %v", err)
		http.Error(w, "Failed to read munros data", http.StatusInternalServerError)
		return
	}

	// Apply filters if provided
	query := r.URL.Query()
	filteredMunros := filterMunros(munros, query)

	writeJSONResponse(w, filteredMunros, http.StatusOK)
}

// Get specific munro by ID
func handleMunroByID(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	munroID := r.PathValue("id")
	if munroID == "" {
		http.Error(w, "Missing munro ID", http.StatusBadRequest)
		return
	}

	munros, err := dataService.ReadMunros()
	if err != nil {
		log.Printf("Error reading munros: %v", err)
		http.Error(w, "Failed to read munros data", http.StatusInternalServerError)
		return
	}

	// Find munro by ID (can be running number or DoBIH number)
	id, _ := strconv.Atoi(munroID)
	for _, munro := range munros {
		if munro.RunningNo == id || munro.DoBIHNumber == id {
			writeJSONResponse(w, munro, http.StatusOK)
			return
		}
	}

	http.Error(w, "Munro not found", http.StatusNotFound)
}

// Legacy endpoint for CSV data
func handleMunrosCSV(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	// Just redirect to JSON endpoint
	handleGetMunros(w, r)
}

// Alias for handleGetMunros
func handleGetAllMunros(w http.ResponseWriter, r *http.Request) {
	handleGetMunros(w, r)
}

// Helper function to filter munros based on query parameters
func filterMunros(munros []model.Munro, query map[string][]string) []model.Munro {
	var filtered []model.Munro

	for _, munro := range munros {
		include := true

		// Filter by classification (munro, top, other)
		if classification := query["classification"]; len(classification) > 0 {
			if !strings.EqualFold(munro.Classification, classification[0]) {
				include = false
			}
		}

		// Filter by minimum height
		if minHeight := query["min_height"]; len(minHeight) > 0 {
			if height, err := strconv.ParseFloat(minHeight[0], 64); err == nil {
				if munro.HeightM < height {
					include = false
				}
			}
		}

		// Filter by SMC section
		if section := query["section"]; len(section) > 0 {
			if !strings.Contains(strings.ToLower(munro.SMCSection), strings.ToLower(section[0])) {
				include = false
			}
		}

		// Filter by name search
		if search := query["search"]; len(search) > 0 {
			if !strings.Contains(strings.ToLower(munro.Name), strings.ToLower(search[0])) {
				include = false
			}
		}

		if include {
			filtered = append(filtered, munro)
		}
	}

	return filtered
}

// ###########################################
// Handling Pages
// ###########################################
func handleIndex(w http.ResponseWriter, r *http.Request) {
	//Server the index.html file
	http.ServeFile(w, r, "./frontend/index.html")
}
