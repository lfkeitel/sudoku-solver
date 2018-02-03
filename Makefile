.PHONY: build-go build-java

build-java:
	javac Solver.java

build-go:
	go build -o bin/sudoku ./...
