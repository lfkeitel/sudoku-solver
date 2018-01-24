package main

import (
	"testing"
)

func BenchmarkEasyPuzzle(b *testing.B) {
	grid := readBoardFromFile("testdata/easy-grid")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		depthFirstSolve(grid)
	}
}

func BenchmarkMediumPuzzle(b *testing.B) {
	grid := readBoardFromFile("testdata/medium-grid")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		depthFirstSolve(grid)
	}
}

func BenchmarkHardPuzzle(b *testing.B) {
	grid := readBoardFromFile("testdata/hard-grid")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		depthFirstSolve(grid)
	}
}

func BenchmarkEvilPuzzle(b *testing.B) {
	grid := readBoardFromFile("testdata/evil-grid")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		depthFirstSolve(grid)
	}
}
