# Sudoku Puzzle Solver

This solver uses depth-first search to solve Sudoku puzzle. Grids are inputted as files
with a grid structure. Examples can be found in the grids directory.

## Install

`go get github.com/lfkeitel/sudoku-solver`

`sudoku-solver -grid $PATH_TO_GRID_FILE`

## Benchmarks

Processor: Intel Core i7-6700K CPU @ 4.00GHz
RAM: 32GB

```
goos: linux
goarch: amd64
pkg: github.com/lfkeitel/sudoku-solver
BenchmarkEasyPuzzle-8              20000             92146 ns/op
BenchmarkMediumPuzzle-8             1000           1533778 ns/op
BenchmarkHardPuzzle-8                 30          47485735 ns/op
BenchmarkEvilPuzzle-8                 50          32876005 ns/op
```
