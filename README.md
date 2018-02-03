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

## Grid Files

Grid files allow the user to input a sudoku grid to solve. The format is very simple:

```
0 6 0 0 2 1 0 0 7
3 1 0 6 0 7 0 0 0
8 0 0 5 0 0 0 0 0
0 0 8 0 1 0 0 0 6
0 2 0 7 5 9 0 3 0
7 0 0 0 8 0 1 0 0
0 0 0 0 0 5 0 0 8
0 0 0 2 0 8 0 9 3
2 0 0 9 7 0 0 4 0
```

It's simply a 9x9 grid of numbers. Pipe separators and dashes can be used make it
look nicer:

```
0 6 0 | 0 2 1 | 0 0 7
3 1 0 | 6 0 7 | 0 0 0
8 0 0 | 5 0 0 | 0 0 0
------+-------+------
0 0 8 | 0 1 0 | 0 0 6
0 2 0 | 7 5 9 | 0 3 0
7 0 0 | 0 8 0 | 1 0 0
------+-------+------
0 0 0 | 0 0 5 | 0 0 8
0 0 0 | 2 0 8 | 0 9 3
2 0 0 | 9 7 0 | 0 4 0
```

Any line that begins with a dash "-" is a comment and ignored. Please see the example
grid files for reference.
