package minesweeper

import (
	"errors"
	"fmt"
	"minesweeper-API/types"
)

// Struct used to handle all the phases of a game.
// This is an immplementation fo the interface types.game_types.GameService
type GameService struct {
	Store types.GameStore
}

// Values used as default.
const (
	defaultRows  = 6
	defaultCols  = 6
	defaultMines = 12
	maxRows      = 30
	maxCols      = 30
)

/*
	Handles all the process related to create a specific game
	Setup the values received as parameters or it will use default values.
 */
func (s *GameService) Create(game *types.Game) error {
	if game.Name == "" {
		return errors.New("no Game name")
	}

	if game.Rows == 0 {
		game.Rows = defaultRows
	}

	if game.Cols == 0 {
		game.Cols = defaultCols
	}

	if game.Mines == 0 {
		game.Mines = defaultMines
	}

	if game.Rows > maxRows {
		game.Rows = maxRows
	}

	if game.Cols > maxCols {
		game.Cols = maxCols
	}

	if game.Mines > (game.Cols * game.Rows) {
		game.Mines = game.Cols * game.Rows
	}

	game.Status = "new"

	err := s.Store.Insert(game)
	return err
}

/*
	This function will handle the process of perform the configuration of a new game previously created.
 */
func (s *GameService) Start(name string) (*types.Game, error) {
	game, err := s.Store.GetByName(name)
	if err != nil {
		return nil, err
	}

	buildBoard(game)

	game.Status = "started"
	err = s.Store.Update(game)
	fmt.Printf("%#v\n", game.Grid)
	return game, err
}
