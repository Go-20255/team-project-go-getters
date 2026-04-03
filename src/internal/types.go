package internal

// TYPES and such
type TileState int


type Tile struct {
	State           TileState
	HasMine         bool
	AdjacentMines int
}

type Controller struct {
	Tiles          [][]Tile
	Width          int
	Height         int
	MineCount      int
	FirstMove      bool
	GameOver       bool
	GameWon        bool
}