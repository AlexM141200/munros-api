package api

import (
	"log"
	"net/http"
  "path/filePath"
  "os"
  "database/sql"

  _ "github.com/mattn/go-sqlite3" 
)

type APIServer struct {
	addr string
}

type Application struct{
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
  dataDir := "./data"
  dbFilename := "munro.db"
  dbPath := filePath.Join(datadir, dbFilename)


  //Create data directory if not exists
  if _, err := os.Stat(dataDir); os.IsNotExists(err) {
          fmt.Printf("Creating Data Directory. %s\n", dataDir)
            if err := os.MkDir(dataDir, 0755); err != nil {
              return fmt.Errorf("Failed to create directory.")
    }
  }

  //Open sqlite database.
  db, err := sql.Open()


  app := &Application{
    DB: db,
  }


	router := http.NewServeMux()

	//GET Requests
	router.HandleFunc("/", handleIndex)
	router.HandleFunc("/munros/{id}", handleMunroByID)
	router.HandleFunc("/munros/", handleGetMunros)
	router.HandleFunc("/munrosCSV/", handleMunrosCSV)
  router.HandleFunc("/getAllMunros/", handleGetAllMunros)

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server running on port %s", s.addr)
	return server.ListenAndServe()
}
