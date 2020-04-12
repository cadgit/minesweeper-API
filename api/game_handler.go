package api

import (
	"encoding/json"
	"minesweeper-API/types"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

/*
title: create game

This function will handle the request associated to the process of create a Game.
It will create the game an store the information.

path: /game
method: POST
responses:
  201: Game created
  400: Invalid json
  500: server error
*/
func (s *Services) createGame(w http.ResponseWriter, r *http.Request) {
	var game types.Game

	log := s.logger.WithFields(logrus.Fields{
		"s":      "game",
		"method": "create",
	})

	errorParseData := json.NewDecoder(r.Body).Decode(&game)

	if errorParseData != nil {
		log.Error(errorParseData)
		ErrInvalidJSON.Send(w)
		return
	}

	errCreateGame := s.GameService.Create(&game)

	if errCreateGame != nil {
		log.WithField("err", errCreateGame).Error("cannot create game")
		ErrInternalServer.Send(w)
		return
	}

	Success(game, http.StatusCreated).Send(w)
}

/*
title: start game

This function will handle the request related to the configuration of a game already created
but it wasn't initiated by the player.

path: /game/{name}/start
method: POST
responses:
  200: OK
  500: server error
*/
func (s *Services) startGame(respWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]

	log := s.logger.WithFields(logrus.Fields{
		"service": "game",
		"method":  "start",
	})

	game, err := s.GameService.Start(name)
	if err != nil {
		log.WithField("err", err).Error("cannot start game")
		ErrInternalServer.Send(respWriter)
		return
	}

	Success(game, http.StatusOK).Send(respWriter)
}
