package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	gridFile string
	debug    bool
	pathOut  string
)

func init() {
	flag.StringVar(&gridFile, "grid", "grid", "Starting grid file")
	flag.BoolVar(&debug, "debug", false, "Enable debug mode")
	flag.StringVar(&pathOut, "po", "", "Output the solve path to FILENAME")
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

	fmt.Printf("\nSolved grid in %s with %d states:\n\n", dur.String(), len(path))
	grid2.print("  ")

	if pathOut != "" {
		pathOutFile, err := os.OpenFile(pathOut, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		pathLen := len(path)
		for i := pathLen - 1; i >= 0; i-- {
			pathOutFile.WriteString("State " + strconv.Itoa(pathLen-i) + ": \n")
			path[i].fprint(pathOutFile, "")
			pathOutFile.Write([]byte{'\n'})
		}

		pathOutFile.Close()
	}
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
