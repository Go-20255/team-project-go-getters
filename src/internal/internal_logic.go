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

func PlaceMines(tiles [][]Tile, width, height, count, avoidX, avoidY int) {
	if count > width*height {
		panic("Trying to initialize more mines than tiles")
	}

	coordinates := make([][]int, (width * height))
	i := 0
	j := 0
	for pair := range coordinates {
		coordinates[pair] = make([]int, 2)
		coordinates[pair][0] = i
		coordinates[pair][1] = j
		i++
		if i >= width {
			j++
			i = 0
		}
	}

	for range count {
		idx := rand.Intn(len(coordinates))
		pair := coordinates[idx]
		tiles[pair[0]][pair[1]].HasMine = true

		coordinates[idx] = coordinates[len(coordinates)-1]
		coordinates = coordinates[:len(coordinates)-1]
	}
}

func CalculateAdjacency(tiles [][]Tile, width, height int) {
	for i := range tiles {
		for j := range tiles[i] {
			if !tiles[i][j].HasMine {
				tiles[i][j].AdjacentMines = CountNeighborMines(tiles, i, j, width, height)
			}
		}
	}
}

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

func FloodReveal(tiles [][]Tile, tx, ty, width, height int) {
}
