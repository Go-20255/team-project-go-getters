package internal

// controller initialization
func NewController(width, height, mineCount int) *Controller {
	return &Controller{
		Width:     width,
		Height:    height,
		MineCount: mineCount,
		FirstMove: true,
	}
}

// reveals the specified tile. If that tile has no neighboring
// mines, moves into the recursive FloodReveal
func (c *Controller) RevealTile(tx, ty int) {
	if c.Tiles[tx][ty].State == 1 {
		panic("Cannot reveal already revealed tile")
	}
	if c.Tiles[tx][ty].HasMine {
		c.GameOver = true
	} else if c.Tiles[tx][ty].AdjacentMines == 0 {
		FloodReveal(c.Tiles, tx, ty, c.Width, c.Height)
		return
	}
	c.Tiles[tx][ty].State = 1
}

// sets the specified tile to the `flagged` state
func (c *Controller) FlagTile(tx, ty int) {
	c.Tiles[tx][ty].State = 2
}

// TODO: Unsure what this is supposed to do
func (c *Controller) Reset() {
}

// is the game over?
func (c *Controller) IsGameOver() bool {
	return c.GameOver
}

// is the game won?
func (c *Controller) IsGameWon() bool {
	return c.GameWon
}

// returns the specified tile
func (c *Controller) GetTile(tx, ty int) Tile {
	return c.Tiles[ty][tx]
}
