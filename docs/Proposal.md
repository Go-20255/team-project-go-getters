# Minesweeper Proposal:
**Team Members:**
* Jack Werremeyer
* Nolan Byer

## 1. Project Summary
Minesweeper is a single-player game where players uncover tiles on a grid while avoiding hidden mines. Players can flag suspected mines and use numerical clues to deduce safe tiles. Each revealed tile either shows a number indicating adjacent mines, is blank, or triggers a game-ending explosion.

Hopefully, by the end of the project, the system will resemble a fully playable Minesweeper experience, where logic and deduction result in fun and satisfying gameplay.

## 2. Typical Use Cases
1. **Playing the Game**
The backend server manages the Minesweeper board state, handling mine placement, tile reveals, and win/loss detection in real time.

2. **GUI Monitoring and Control**
Users interact with the game using a GUI client to visualize the board, reveal or flag tiles, restart the game, and adjust parameters like grid size or mine density.

## 3. Intended Design and Components
The project will be organized into multiple components:

### Core Modules
1. minesweeper/internal
   * Contains the core game logic like mine placement, tile state, adjacency calculation)
2. minesweeper/controller
   * Controls player actions and keeps the current state of the board

### Simulation/System Modules
3. game/world
   * The minefield grid that handles tile generation, mine seeding, reveal propagation, and flag placement
4. game/server
   * Multi-threaded Go server that runs the game loop and processes player input
5. client/gui
   * GUI for visualizing the minefield and controlling the game

## 4. Testing Strategy
Testing will largely focus on ensuring the core game logic functions as intended — particularly mine placement, tile reveal propagation, and win/loss detection. We'll also be testing the GUI for visual bugs.

## 5. Minimal Viable Product
The MVP will include:
* A functional minefield grid with randomized mine placement
* Tile reveal logic
* Basic flagging and win/loss detection
* A command-line or minimal visualization of the board
* Necessary unit tests for the above

## 6. Stretch Goals
* A polished GUI visualization (classic Minesweeper?)
* Difficulty presets (Beginner, Intermediate, Expert)
* A custom game mode with configurable grid size and mine count
* A leaderboard or timer for tracking best times
* Chord-clicking and other advanced player interactions

## 7. Expected Functionality by Checkpoint
By the checkpoint, the project is expected to include:
* Working core game logic (mine placement, reveal, flagging)
* Basic structure of the game board
* Unit tests
* Basic scaffolding for the game server