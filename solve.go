package main

import (
	"errors"
)

func depthFirstSolve(grid board) (board, error) {
	row, col := grid.findNextUnassigned()
	if row == -1 { // Grid is solved
		return grid, nil
	}

	for i := 1; i <= 9; i++ {
		if grid.isValidMove(i, row, col) {
			debugPrintf("Trying Row: %d, Col: %d, Val %d\n", row, col, i)
			grid[row][col] = i

			if grid, err := depthFirstSolve(grid); err == nil {
				return grid, nil // Grid is solved
			}

			grid[row][col] = unassigned
		}
	}

	return grid, errors.New("Not solved")
}
