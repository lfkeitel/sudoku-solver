package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	unassigned = 0

	gridHeight = 9
	gridWidth  = 9
)

type board [gridHeight][gridWidth]int

func readBoardFromFile(filepath string) board {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	var grid board

	reader := bytes.NewBuffer(file)

	lineNum := 0
	rowNum := 0
	for {
		lineNum++

		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatalln(err)
			}
			break
		}

		line = bytes.TrimSpace(line)

		if len(line) == 0 || line[0] == '-' { // Skip empty lines or lines starting with a dash (grid separator/comments)
			continue
		}

		cells := bytes.Split(line, []byte{' '})
		if len(cells) != 9 && len(cells) != 11 {
			log.Fatalf("Error in grid file on line %d\n", lineNum)
		}

		cellNum := 0
		for _, cell := range cells {
			if len(cell) != 1 {
				continue
			}

			if cell[0] == '|' {
				continue
			}
			cellValue := int(cell[0]) - '0'
			if cellValue < 0 || cellValue > 9 {
				log.Fatalf("Invalid cell value in grid on line %d\n", lineNum)
			}
			grid[rowNum][cellNum] = cellValue
			cellNum++
		}
		rowNum++
	}

	return grid
}

func (b board) print(indent string) {
	b.fprint(os.Stdout, indent)
}

func (b board) fprint(w io.Writer, indent string) {
	for i := 0; i < gridHeight; i++ {
		fmt.Fprint(w, indent)

		for j := 0; j < gridWidth; j++ {
			fmt.Fprintf(w, "%d ", b[i][j])

			if (j+1)%3 == 0 && j < gridWidth-1 {
				fmt.Fprint(w, "| ")
			}
		}
		fmt.Fprintln(w, "")

		if (i+1)%3 == 0 && i < gridHeight-1 {
			fmt.Fprint(w, indent)
			fmt.Fprintln(w, "------+-------+------")
		}
	}
}

// findNextUnassigned will return the first cell starting at 0,0 going by row them column that has the value 0.
func (b board) findNextUnassigned() (int, int) {
	for i := 0; i < gridHeight; i++ {
		for j := 0; j < gridWidth; j++ {
			if b[i][j] == unassigned {
				return i, j
			}
		}
	}
	return -1, -1
}

func (b board) solved() bool {
	row, _ := b.findNextUnassigned()
	return row == -1
}

// isValidMove returns if the cell at row, col could have the value val.
func (b board) isValidMove(val, row, col int) bool {
	return b.isValidRow(val, row) &&
		b.isValidCol(val, col) &&
		b.isValidNonet(val, row, col)
}

func (b board) isValidRow(val, row int) bool {
	for _, cell := range b[row] {
		if cell == val {
			return false
		}
	}
	return true
}

func (b board) isValidCol(val, col int) bool {
	for _, cell := range b {
		if cell[col] == val {
			return false
		}
	}
	return true
}

func (b board) isValidNonet(val, row, col int) bool {
	nonetRowStart := row - (row % 3)
	nonetColStart := col - (col % 3)

	debugPrintf("nonetRow: %d, nonetCol: %d\n", nonetRowStart, nonetColStart)

	for i := nonetRowStart; i < nonetRowStart+3; i++ {
		for j := nonetColStart; j < nonetColStart+3; j++ {
			debugPrintf("%d %d\n", val, b[i][j])
			if b[i][j] == val {
				return false
			}
		}
	}
	return true
}
