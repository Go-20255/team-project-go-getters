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
	coordinates := make([][2]int, 0, (width*height)-1)
	for y := range height {
		for x := range width {
			if y == avoidY && x == avoidX {
				continue
			}
			coordinates = append(coordinates, [2]int{y, x})
		}
	}

	// initialize a mine at a random coordinate and remove that
	// from the list of available coordinates, repeat `count` times
	for range count {
		idx := rand.Intn(len(coordinates))
		coordinate := coordinates[idx]
		tiles[coordinate[0]][coordinate[1]].HasMine = true

		coordinates[idx] = coordinates[len(coordinates)-1]
		coordinates = coordinates[:len(coordinates)-1]
	}
}

// initialize each tile's `AdjacentMines` field
func CalculateAdjacency(tiles [][]Tile, width, height int) {
	for y := range tiles {
		for x := range tiles[y] {
			if !tiles[y][x].HasMine { // skip tiles with mines
				tiles[y][x].AdjacentMines = CountNeighborMines(tiles, x, y, width, height)
			}
		}
	}
}

// helper function for CalculateAdjacency, called for each tile.
// Counts mines in the surrounding tiles, returns the total count.
// Uses `min` and `max` to avoid out of bounds indexing
func CountNeighborMines(tiles [][]Tile, tx, ty, width, height int) int {
	mineCount := 0

	for y := max(0, ty-1); y <= min(height-1, ty+1); y++ {
		for x := max(0, tx-1); x <= min(width-1, tx+1); x++ {
			if tiles[y][x].HasMine {
				mineCount++
			}
		}
	}

	return mineCount
}

// recursively reveal all adjacent tiles around tiles with 0 adjacent mines
func FloodReveal(tiles [][]Tile, tx, ty, width, height int) {
	if tiles[ty][tx].State != TileHidden { // avoid infinite loops of revealing already revealed tiles
		return
	}

	tiles[ty][tx].State = TileRevealed // reveal this tile

	if tiles[ty][tx].AdjacentMines == 0 {
		for y := max(0, ty-1); y <= min(height-1, ty+1); y++ {
			for x := max(0, tx-1); x <= min(width-1, tx+1); x++ {
				// this will iterate over the current tile, but that will be caught by the above return
				FloodReveal(tiles, x, y, width, height)
			}
		}
	}
}
