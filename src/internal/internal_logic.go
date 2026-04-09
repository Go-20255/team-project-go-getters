package internal

import (
	"math/rand"
)

// Tile state constants
const (
	TileHidden TileState = iota
	TileRevealed
	TileFlagged
)

// initialize `count` mines on a new board at random tiles,
// with the exception of tile (`avoidX`, `avoidY`)
func PlaceMines(tiles [][]Tile, width, height, count, avoidX, avoidY int) {
	if count > width*height-1 {
		panic("Trying to initialize more mines than tiles")
	}

	// a 2d array of pairs of integers, representing tile coordinates
	coordinates := make([][]int, (width*height)-1)
	i := 0
	j := 0
	for pair := range coordinates {
		if !(i == avoidX && j == avoidY) { // skip the avoided tile
			coordinates[pair] = make([]int, 2)
			coordinates[pair][0] = i
			coordinates[pair][1] = j
		}
		i++
		if i >= width {
			j++
			i = 0
		}
	}

	// initialize a mine at a random coordinate and remove that
	// from the list of available coordinates, repeat `count` times
	for range count {
		idx := rand.Intn(len(coordinates))
		pair := coordinates[idx]
		tiles[pair[0]][pair[1]].HasMine = true

		coordinates[idx] = coordinates[len(coordinates)-1]
		coordinates = coordinates[:len(coordinates)-1]
	}
}

// initialize each tile's `AdjacentMines` field
func CalculateAdjacency(tiles [][]Tile, width, height int) {
	for i := range tiles {
		for j := range tiles[i] {
			if !tiles[i][j].HasMine { // skip tiles with mines
				tiles[i][j].AdjacentMines = CountNeighborMines(tiles, i, j, width, height)
			}
		}
	}
}

// helper function for CalculateAdjacency, called for each tile.
// Counts mines in the surrounding tiles, returns the total count.
// Uses `min` and `max` to avoid out of bounds indexing
func CountNeighborMines(tiles [][]Tile, tx, ty, width, height int) int {
	mineCount := 0

	for i := max(0, tx-1); i <= min(width-1, tx+1); i++ {
		for j := max(0, ty-1); j <= min(height-1, ty+1); j++ {
			if tiles[i][j].HasMine {
				mineCount++
			}
		}
	}

	return mineCount
}

// recursively reveal all adjacent tiles around tiles with 0 adjacent mines
func FloodReveal(tiles [][]Tile, tx, ty, width, height int) {
	if tiles[tx][ty].State == 1 { // avoid infinite loops of revealing already revealed tiles
		return
	}
	tiles[tx][ty].State = 1 // reveal this tile
	if tiles[tx][ty].AdjacentMines == 0 {
		for i := max(0, tx-1); i <= min(width-1, tx+1); i++ {
			for j := max(0, tx-1); j <= min(height-1, ty+1); j++ {
				// this will iterate over the current tile, but that will be caught by the above return
				FloodReveal(tiles, i, j, width, height)
			}
		}
	}
}
