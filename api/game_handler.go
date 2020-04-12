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

/*
title: cell click

This function will handle the request related to perform the action of select a
specific cell in the board.

path: /game/{name}/click
method: POST
responses:
  200: OK
  400: Invalid json
  500: server error
 */
func (s *Services) clickCell(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	log := s.logger.WithFields(logrus.Fields{
		"service": "game",
		"method":  "click",
	})

	var cellPos struct {
		Row int `json:"row"`
		Col int `json:"col"`
	}

	if err := json.NewDecoder(r.Body).Decode(&cellPos); err != nil {
		log.Error(err)
		ErrInvalidJSON.Send(w)
		return
	}

	game, err := s.GameService.Click(name, cellPos.Row, cellPos.Col)
	if err != nil {
		log.WithField("err", err).Error("cannot click cell")
		ErrInternalServer.Send(w)
		return
	}
	cell := game.Grid[cellPos.Row][cellPos.Col]

	if game.Status != "over" && game.Status != "won" {
		game.Grid = nil
	}

	var result struct {
		Cell types.Cell
		Game types.Game
	}

	result.Cell = cell
	result.Game = *game

	Success(&result, http.StatusOK).Send(w)
}
