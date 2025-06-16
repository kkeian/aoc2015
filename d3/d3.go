package main

import (
	"bufio"
	"fmt"
	"os"
)

type direction int

const (
	Up direction = iota
	Down
	Left
	Right
)

var move = map[byte]direction{
	'^': Up,
	'v': Down,
	'>': Right,
	'<': Left,
}

type houseCoord struct {
	X, Y int
}

func main() {
	// start at 1 because Santa always delivers to the
	// house at the location he starts
	uniqueHouses := 1

	content, err := os.ReadFile("input")
	if err != nil {
		os.Exit(1)
	}

	startPos := houseCoord{0, 0}
	currPos := startPos
	housesVisted := make(map[houseCoord]bool)
	housesVisted[startPos] = true

	buf := bufio.NewWriter(os.Stdout)
	lineOut := fmt.Sprintf("1 - %c (%d, %d)\n", ' ', currPos.X, currPos.Y)
	_, err = buf.WriteString(lineOut)
	if err != nil {
		os.Stderr.WriteString("Error writing line\n")
	}

	for i, c := range content {
		var nextPos houseCoord
		switch move[c] {
		case Up:
			nextPos = houseCoord{currPos.X, currPos.Y + 1}
		case Down:
			nextPos = houseCoord{currPos.X, currPos.Y - 1}
		case Left:
			nextPos = houseCoord{currPos.X - 1, currPos.Y}
		case Right:
			nextPos = houseCoord{currPos.X + 1, currPos.Y}
		}

		if !housesVisted[nextPos] {
			housesVisted[nextPos] = true
			uniqueHouses++
		}

		lineOut = fmt.Sprintf("%d - %c (%d, %d)\n", i+1, c, nextPos.X, nextPos.Y)
		_, err = buf.WriteString(lineOut)
		if err != nil {
			os.Stderr.WriteString("Error writing line\n")
		}

		currPos = nextPos
	}

	lineOut = fmt.Sprintf("Number of unique houses presents delivered at: %d\n", uniqueHouses)
	_, err = buf.WriteString(lineOut)
	if err != nil {
		os.Stderr.WriteString("Error writing line\n")
	}
	err = buf.Flush()
	if err != nil {
		os.Stderr.WriteString("Error flushing buffer\n")
	}
}
