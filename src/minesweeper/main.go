package main

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"

	"github.com/Go-20255/team-project-go-getters/src/internal"
	"github.com/Go-20255/team-project-go-getters/src/game_graphics"
)

const (
    boardWidth  = 16
    boardHeight = 16
    mineCount   = 40
    tileSize    = 32
)

func main() {
    ctrl := internal.NewController(boardWidth, boardHeight, mineCount)

    bg := game_graphics.NewBoardGen(ctrl, tileSize)
    bg.GenerateGrid()

    g := game_graphics.NewGUI(bg, boardWidth*tileSize, boardHeight*tileSize)

    gl := game_graphics.NewGameLoop(bg, g)
    gl.Start()
    defer gl.Stop()

    ebiten.SetWindowSize(boardWidth*tileSize, boardHeight*tileSize)
    ebiten.SetWindowTitle("Minesweeper")

    if err := ebiten.RunGame(g); err != nil {
        log.Fatal(err)
    }
}