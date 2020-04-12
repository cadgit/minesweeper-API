package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"

	"minesweeper-API/types"
)

type Services struct {
	logger      *logrus.Logger
	GameService types.GameService
}

func Start(log *logrus.Logger) error {
	services := Services{
		logger: log,
	}

	// API Routes
	r := Router(&services)

	// Middleware
	n := negroni.Classic()
	n.UseHandler(r)

	//Run Server
	log.Infoln("Server running on port :8080")
	http.ListenAndServe(":8080", n)
	return nil
}

func Router(services *Services) *mux.Router {
	// API Routes
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", services.healthcheck).Methods("GET")
	r.HandleFunc("/game", services.createGame).Methods("POST")
	r.HandleFunc("/game/{name}/start", services.startGame).Methods("POST")
	return r
}
