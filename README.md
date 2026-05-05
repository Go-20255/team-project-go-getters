# Go Getters Semester Project: Minesweeper
**Team Members:**
* Jack Werremeyer
* Nolan Byer

## Game Summary
Minesweeper is a single-player game where players uncover tiles on a grid while avoiding hidden mines. Hidden tiles can be revealed by left clicking on them. Each revealed tile either shows a number indicating adjacent mines in the surrounding 3x3 grid, is blank to indicate zero adjacent mines, or triggers a game-ending explosion if a mine is uncovered. The goal of the game is to reveal all non-mine tiles.
Convenience features on top of the above baseline include:
* Blank tiles will automatically reveal all adjacent tiles when revealed, as the safety of these tiles is a trivial deduction.
* A timer is displayed at the top left, for players to race their own previous attempts if they so choose.
* Hidden tiles can be 'flagged' with right click, leaving an indicator that the player suspects that tile to be a mine.
* A number of flags available is displayed to the player, corresponding to the number of hidden mines. These decrement as flags are placed, indicating to the player the number of mines remaining if their deductions are assumed to be correct.

## Project Structure
The project is laid out as follows:
```
src
├── game_graphics      // front end
│   ├── assets
│   │    └── [images…]      // all game sprites are stored here
│   ├── game_loop.go        // listens for input and makes calls to backend
│   ├── board_gen.go        // initializes the board
│   ├── gui.go              // renders the game window
│   └── types.go            // type definitions for front end
│
├── internal                // back end
│    ├── controller.go      // provides methods for frontend to call, houses the model
│    ├── internal_logic.go  // algorithms for game actions
│    └── types.go           // type definitions for back end
│
└── minesweeper
    └── main.go             // program entrypoint, initializes board parameters
```

## Usage
To run the program, simply use `go run src/minesweeper/main.go` from the project root. Alternatively, `go run main.go` from `src/minesweeper` also works.
Doing so will open up a new game window, with timer started and all tiles hidden. Click a tile to begin playing, and have fun!