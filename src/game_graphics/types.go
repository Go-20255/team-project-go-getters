package game_graphics

import (
	//"fmt"
	//"image"
	"sync"

	//"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/colorm"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/Go-20255/team-project-go-getters/src/internal"
)


// USE FOR TYPES AND STRUCTS

//GUI types


type GUI struct {
	BoardGen *BoardGen
	Width    int
	Height   int
}


//game_loop types

type InputEvent struct {
    Type InputType
    X, Y int
}

type GameLoop struct {
    BoardGen  *BoardGen
    GUI       *GUI
    InputChan chan InputEvent
    QuitChan  chan struct{}
    wg        sync.WaitGroup
    mu        sync.Mutex
}

type InputType int



//Board Gen types

type BoardGen struct {
	Controller *internal.Controller
	TileSize   int
}