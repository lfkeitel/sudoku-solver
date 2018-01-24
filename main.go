package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var (
	gridFile string
	debug    bool
)

func init() {
	flag.StringVar(&gridFile, "grid", "grid", "Starting grid file")
	flag.BoolVar(&debug, "debug", false, "Enable debug mode")
}

func main() {
	flag.Parse()

	grid := readBoardFromFile(gridFile)

	fmt.Print("Input grid:\n\n")
	grid.print("  ")

	start := time.Now()
	grid2, err := depthFirstSolve(grid)
	dur := time.Now().Sub(start)
	if err != nil {
		log.Fatal("Could not find a solution. Are you sure it's solvable?")
	}

	fmt.Printf("\nSolved grid in %s:\n\n", dur.String())
	grid2.print("  ")
}

func debugPrintln(msg ...interface{}) {
	if debug {
		log.Println(msg...)
	}
}

func debugPrintf(msg string, a ...interface{}) {
	if debug {
		log.Printf(msg, a...)
	}
}
