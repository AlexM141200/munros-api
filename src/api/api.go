package api

import (
	"database/sql"
	"log"
	"net/http"

	"context"

	_ "github.com/mattn/go-sqlite3"

	"github.com/AlexM141200/munros-api/src/handlers"
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
func (s *APIServer) Run(ctx context.Context) error {

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
		DB: nil,
	}

	router := http.NewServeMux()

	// API Routes
	handlers.SetupMunroRoutes(router)

	// Frontend Routes
	handlers.SetupFrontendRoutes(router)

	fs := http.FileServer(http.Dir("./munromark/build/client/"))
	router.Handle("/assets/", http.StripPrefix("/assets/", fs))

	_ = app

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server running on port %s", s.addr)
	return server.ListenAndServe()
}
