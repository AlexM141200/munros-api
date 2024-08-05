package api

import (
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

// Run Function of API Server
func (s *APIServer) Run() error {

	router := http.NewServeMux()

	//GET Requests
	router.HandleFunc("/munros/{id}", handleMunroByID)
	router.HandleFunc("/munros", handleGetMunros)

	//POST Requests
	router.HandleFunc("POST /munros/{id}", handlePostMunro)

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server running on port %s", s.addr)
	return server.ListenAndServe()
}
