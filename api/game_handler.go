package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"minesweeper-API/types"
	"net/http"
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
		"s": "game",
		"method":  "create",
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
