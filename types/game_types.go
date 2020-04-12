package types

type Game struct {
	Name string `json:"name"`
}

type GameService interface {
	Create(game *Game) error
	Start(name string) (*Game, error)
	Click(name string, i, j int) (*Game, error)
}
