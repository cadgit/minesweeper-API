package minesweeper

import (
	"minesweeper-API/tests"
	"minesweeper-API/types"
	"reflect"
	"testing"
)

/*
	This test case performs a validation of the logic
	associated with creation of a new game.
 */
func TestCreateGame(t *testing.T) {
	s := GameService{
		Store: &tests.MockGameStore{
			OnInsert: func(game *types.Game) error {
				return nil
			},
		},
	}
	game := &types.Game{
		Name:  "mockGame",
		Cols:  10,
		Rows:  11,
		Mines: 12,
	}

	if err := s.Create(game); err != nil {
		t.Fatal(err)
	}

	if game.Cols != 10 {
		t.Errorf("unexpected cols. want=10, got %d", game.Cols)
	}
	if game.Rows != 11 {
		t.Errorf("unexpected rows. want=11, got %d", game.Rows)
	}
	if game.Mines != 12 {
		t.Errorf("unexpected mines. want=12, got %d", game.Mines)
	}
	if game.Status != "new" {
		t.Errorf("unexpected status. want='new', got %s", game.Status)
	}
}

/*
	This test case perform a validation of the logic related to
	click a specific cell
 */
func TestClickCell(t *testing.T) {
	s := GameService{
		Store: &tests.MockGameStore{
			OnGetByName: func(name string) (*types.Game, error) {
				grid := []types.CellGrid{
					types.CellGrid{
						types.Cell{Mine: false, Clicked: false, Value: 1},
						types.Cell{Mine: true, Clicked: false, Value: 0},
					},
					types.CellGrid{
						types.Cell{Mine: false, Clicked: false, Value: 1},
						types.Cell{Mine: false, Clicked: false, Value: 1},
					},
				}
				return &types.Game{
					Name:   name,
					Cols:   2,
					Rows:   2,
					Mines:  1,
					Status: "started",
					Grid:   grid,
				}, nil
			},
			OnUpdate: func(game *types.Game) error {
				return nil
			},
		},
	}

	game, err := s.Click("test", 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	if game.Status != "started" {
		t.Errorf("unexpected status. want='started', got %s", game.Status)
	}

	expected := []types.CellGrid{
		types.CellGrid{
			types.Cell{Mine: false, Clicked: true, Value: 1},
			types.Cell{Mine: true, Clicked: false, Value: 0},
		},
		types.CellGrid{
			types.Cell{Mine: false, Clicked: false, Value: 1},
			types.Cell{Mine: false, Clicked: false, Value: 1},
		},
	}
	if !reflect.DeepEqual(game.Grid, expected) {
		t.Errorf("unexpected grid. want=%v, got=%v", expected, game.Grid)
	}
}

/*
	This test case check if the logic associated to the situation
	where a player clicked all the cell without a mine.

 */
func TestClickCell_Won(t *testing.T) {
	s := GameService{
		Store: &tests.MockGameStore{
			OnGetByName: func(name string) (*types.Game, error) {
				grid := []types.CellGrid{
					types.CellGrid{
						types.Cell{Mine: false, Clicked: true, Value: 1},
						types.Cell{Mine: true, Clicked: false, Value: 0},
					},
					types.CellGrid{
						types.Cell{Mine: false, Clicked: true, Value: 1},
						types.Cell{Mine: false, Clicked: false, Value: 1},
					},
				}
				return &types.Game{
					Name:   name,
					Cols:   2,
					Rows:   2,
					Mines:  1,
					Clicks: 2,
					Status: "started",
					Grid:   grid,
				}, nil
			},
			OnUpdate: func(game *types.Game) error {
				return nil
			},
		},
	}

	game, err := s.Click("test", 1, 1)
	if err != nil {
		t.Fatal(err)
	}

	if game.Status != "won" {
		t.Errorf("unexpected status. want='won', got %s", game.Status)
	}

	expected := []types.CellGrid{
		types.CellGrid{
			types.Cell{Mine: false, Clicked: true, Value: 1},
			types.Cell{Mine: true, Clicked: false, Value: 0},
		},
		types.CellGrid{
			types.Cell{Mine: false, Clicked: true, Value: 1},
			types.Cell{Mine: false, Clicked: true, Value: 1},
		},
	}
	if !reflect.DeepEqual(game.Grid, expected) {
		t.Errorf("unexpected grid. want=%v, got=%v", expected, game.Grid)
	}
}
