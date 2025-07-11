package api

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type APIServer struct {
	addr string
}

type Application struct {
	DB *sql.DB
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

// Run Function of API Server
func (s *APIServer) Run() error {

	//Data directory
	/*
	  dataDir := "./data"
	  dbFilename := "munro.db"
	*/

	//Open sqlite database.
	/*
	  db, err := sql.Open()
	  if err != nil {
	  panic(err)
	  }

	*/

	app := &Application{
		DB: nil, // Will be initialized when we switch to database
	}

	router := http.NewServeMux()

	// API Routes
	router.HandleFunc("/api/munros", handleGetMunros)
	router.HandleFunc("/api/munros/{id}", handleMunroByID)
	router.HandleFunc("/api/munros/csv", handleMunrosCSV)
	router.HandleFunc("/api/munros/all", handleGetAllMunros)

	// Frontend Routes
	router.HandleFunc("/", handleIndex)

	// Static file server for frontend assets
	fs := http.FileServer(http.Dir("./munromark/build/client/"))
	router.Handle("/assets/", http.StripPrefix("/assets/", fs))

	_ = app // Suppress unused variable warning

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server running on port %s", s.addr)
	return server.ListenAndServe()
}
