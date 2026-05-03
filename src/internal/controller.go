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
	if c.Tiles[ty][tx].State == TileRevealed {
		return // do nothing, tile is already revealed
	}
	if c.Tiles[ty][tx].HasMine {
		c.GameOver = true
	} else if c.Tiles[ty][tx].AdjacentMines == 0 {
		FloodReveal(c.Tiles, tx, ty, c.Width, c.Height)
		c.CheckGameWon()
		return
	}
	c.Tiles[ty][tx].State = TileRevealed
	c.CheckGameWon()
}

// small helper for RevealTile, determines if game is won
// by comparing count of flagged + hidden tiles to total
// mine count
func (c *Controller) CheckGameWon() {
	count := 0
	for _, tileCol := range c.Tiles {
		for _, tile := range tileCol {
			if tile.State != TileRevealed {
				count++
			}
		}
	}
	if count == c.MineCount {
		c.GameWon = true
	}
}

// sets the specified tile to the `flagged` state,
// OR un-flags an already flagged tile.
func (c *Controller) FlagTile(tx, ty int) {
	switch c.Tiles[ty][tx].State {
	case TileHidden:
		c.Tiles[ty][tx].State = TileFlagged
	case TileFlagged:
		c.Tiles[ty][tx].State = TileHidden
	}
}

// Resets all values of the controller.
// NOTE: `GenerateGrid()` has to be called after this to remake Tiles
func (c *Controller) Reset() {
	c.GameOver = false
	c.GameWon = false
	c.FirstMove = true
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
