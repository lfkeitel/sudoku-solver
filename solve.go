package main

import (
	"errors"
)

var path []board

func depthFirstSolve(grid board) (board, error) {
	path = []board{}
	return doDepthFirstSolve(grid)
}

func doDepthFirstSolve(grid board) (board, error) {
	row, col := grid.findNextUnassigned()
	if row == -1 { // Grid is solved
		return grid, nil
	}

	for i := 1; i <= 9; i++ {
		if grid.isValidMove(i, row, col) {
			debugPrintf("Trying Row: %d, Col: %d, Val %d\n", row, col, i)
			grid[row][col] = i

			if grid2, err := depthFirstSolve(grid); err == nil {
				path = append(path, grid)
				return grid2, nil // Grid is solved
			}

			grid[row][col] = unassigned
		}
	}

	return grid, errors.New("Not solved")
}
