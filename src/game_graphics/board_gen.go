package game_graphics

import (
	"github.com/Go-20255/team-project-go-getters/src/internal"
)


func NewBoardGen(ctrl *internal.Controller, tileSize int) *BoardGen {
	return &BoardGen{
		Controller: ctrl,
		TileSize:   tileSize,
	}
}

func (bg *BoardGen) GenerateGrid() {
	bg.Controller.Tiles = make([][]internal.Tile, bg.Controller.Height)
	for y := range bg.Controller.Tiles {
		bg.Controller.Tiles[y] = make([]internal.Tile, bg.Controller.Width)
	}
}

func (bg *BoardGen) SeedMines(avoidX, avoidY int) {
	internal.PlaceMines(
		bg.Controller.Tiles,
		bg.Controller.Width,
		bg.Controller.Height,
		bg.Controller.MineCount,
		avoidX, avoidY,
	)
	internal.CalculateAdjacency(
		bg.Controller.Tiles,
		bg.Controller.Width,
		bg.Controller.Height,
	)
	bg.Controller.FirstMove = false
}

func (bg *BoardGen) PropagateReveal(tx, ty int) {
	if bg.Controller.IsGameOver() || bg.Controller.IsGameWon() {
		return
	}
	if bg.Controller.FirstMove {
		bg.SeedMines(tx, ty)
	}
	bg.Controller.RevealTile(tx, ty)
}

func (bg *BoardGen) PlaceFlag(tx, ty int) {
	if bg.Controller.IsGameOver() || bg.Controller.IsGameWon() {
		return
	}
	bg.Controller.FlagTile(tx, ty)
}

func (bg *BoardGen) TileToPixel(tx, ty int) (int, int) {
	return tx * bg.TileSize, ty * bg.TileSize
}

func (bg *BoardGen) PixelToTile(px, py int) (int, int) {
	return px / bg.TileSize, py / bg.TileSize
}

func (bg *BoardGen) InBounds(tx, ty int) bool {
	return tx >= 0 && ty >= 0 && tx < bg.Controller.Width && ty < bg.Controller.Height
}
