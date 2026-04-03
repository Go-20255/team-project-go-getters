package internal


import (
)

func NewController(width, height, mineCount int) *Controller {
    return &Controller{
        Width:     width,
        Height:    height,
        MineCount: mineCount,
        FirstMove: true,
    }
}

func (c *Controller) RevealTile(tx, ty int) {
}

func (c *Controller) FlagTile(tx, ty int) {
}

func (c *Controller) Reset() {
}

func (c *Controller) IsGameOver() bool {
    return c.GameOver
}

func (c *Controller) IsGameWon() bool {
    return c.GameWon
}

func (c *Controller) GetTile(tx, ty int) Tile {
    return c.Tiles[ty][tx]
}
