package types

/*
	Represent the information that a Cell of the board
	will contain.
 */
type Cell struct {
	Mine    bool `json:"mine"`
	Clicked bool `json:"clicked"`
	Value   int  `json:"value"`
}

/*
	Represent a group of cells that belongs to the board.
 */
type CellGrid []Cell

/*
	Represent the information associated to a game.
 */
type Game struct {
	Name   string     `json:"name"`
	Rows   int        `json:"rows"`
	Cols   int        `json:"cols"`
	Mines  int        `json:"mines"`
	Status string     `json:"status"`
	Grid   []CellGrid `json:"grid,omitempty"`
	Clicks int        `json:"clicks"`
}

/*
	Defines the responsibility of a service
	that it will provide support to the process of handle the different
	phases of a game.
 */
type GameService interface {
	Create(game *Game) error
	Start(name string) (*Game, error)
	Click(name string, i, j int) (*Game, error)
}

/*
	Defines the responsibility of a service that
	it will provide support to the process of handle
	the persistence of the data related to a game.
 */
type GameStore interface {
	Insert(game *Game) error
	Update(game *Game) error
	GetByName(name string) (*Game, error)
}
