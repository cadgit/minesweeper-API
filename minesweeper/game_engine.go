package minesweeper

import (
	"math/rand"
	"minesweeper-API/types"
	"time"
)

/*
	Setup the seed in order to generate random values used in the process of create the mines.
*/
func init() {
	rand.Seed(time.Now().Unix())
}

/*
	It used the information provided at the moment of setup the game in order to generate
	the board related ot the new game.
*/
func buildBoard(game *types.Game) {
	numCells := game.Cols * game.Rows
	cells := make(types.CellGrid, numCells)

	// Randomly set mines
	i := 0
	for i < game.Mines {
		idx := rand.Intn(numCells)
		if !cells[idx].Mine {
			cells[idx].Mine = true
			i++
		}
	}

	game.Grid = make([]types.CellGrid, game.Rows)
	for row := range game.Grid {
		game.Grid[row] = cells[(game.Cols * row):(game.Cols * (row + 1))]
	}

	// Set cell values
	for i, row := range game.Grid {
		for j, cell := range row {
			if cell.Mine {
				setAdjacentValues(game, i, j)
			}
		}
	}
}

/*
	This function will setup the value showed in each cell related to the
	number of mines adjacent to that particular cell.
*/
func setAdjacentValues(game *types.Game, i, j int) {
	for z := i - 1; z < i+2; z++ {
		if z < 0 || z > game.Rows-1 {
			continue
		}
		for w := j - 1; w < j+2; w++ {
			if w < 0 || w > game.Cols-1 {
				continue
			}
			if z == i && w == j {
				continue
			}
			game.Grid[z][w].Value++
		}
	}
}
