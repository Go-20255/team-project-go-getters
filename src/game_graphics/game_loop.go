package game_graphics

import ()


const (
    LeftClick InputType = iota
    RightClick
    Restart
    Quit
)


func NewGameLoop(bg *BoardGen, g *GUI) *GameLoop {
    return &GameLoop{
        BoardGen:  bg,
        GUI:       g,
        InputChan: make(chan InputEvent, 16),
        QuitChan:  make(chan struct{}),
    }
}

func (gl *GameLoop) Start() {
    gl.wg.Add(1)
    go gl.run()
}

func (gl *GameLoop) Stop() {
    close(gl.QuitChan)
    gl.wg.Wait()
}

func (gl *GameLoop) SendInput(e InputEvent) {
    select {
    case gl.InputChan <- e:
    case <-gl.QuitChan:
    }
}

func (gl *GameLoop) run() {
    defer gl.wg.Done()
    for {
        select {
        case <-gl.QuitChan:
            return
        case e := <-gl.InputChan:
            gl.processInput(e)
        }
    }
}

func (gl *GameLoop) processInput(e InputEvent) {
    gl.mu.Lock()
    defer gl.mu.Unlock()

    switch e.Type {
    case LeftClick:
        tx, ty := gl.BoardGen.PixelToTile(e.X, e.Y)
        if gl.BoardGen.InBounds(tx, ty) {
            gl.BoardGen.PropagateReveal(tx, ty)
        }

    case RightClick:
        tx, ty := gl.BoardGen.PixelToTile(e.X, e.Y)
        if gl.BoardGen.InBounds(tx, ty) {
            gl.BoardGen.PlaceFlag(tx, ty)
        }

    case Restart:
        gl.BoardGen.Reset()

    case Quit:
        close(gl.QuitChan)
    }
}