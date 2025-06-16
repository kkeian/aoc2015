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

// bit pack to make keys for houses Visited map
// X = HO 32 bits
// Y = LO 32 bits
// helpers below .X() and .Y()
type houseCoord uint64

func main() {
	// start at 1 because Santa always delivers to the
	// house at the location he starts
	uniqueHouses := 1

	content, err := os.ReadFile("input")
	if err != nil {
		os.Exit(1)
	}

	startPos := houseCoord(0)
	currPos := startPos
	housesVisted := make(map[houseCoord]bool)
	housesVisted[startPos] = true

	buf := bufio.NewWriter(os.Stdout)
	lineOut := fmt.Sprintf("1 - %c (%d, %d)\n", ' ', currPos.X(), currPos.Y())
	_, err = buf.WriteString(lineOut)
	if err != nil {
		os.Stderr.WriteString("Error writing line\n")
	}

	for i, c := range content {
		var nextPos houseCoord
		x, y := currPos.X(), currPos.Y()
		switch move[c] {
		case Up:
			y++
			nextPos = makeHouseCoord(x, y)
		case Down:
			y--
			nextPos = makeHouseCoord(x, y)
		case Left:
			x--
			nextPos = makeHouseCoord(x, y)
		case Right:
			x++
			nextPos = makeHouseCoord(x, y)
		}

		if !housesVisted[nextPos] {
			housesVisted[nextPos] = true
			uniqueHouses++
		}

		lineOut = fmt.Sprintf("%d - %c (%v, %v)\n", i+1, c, nextPos.X(), nextPos.Y())

		_, err = buf.WriteString(lineOut)
		if err != nil {
			os.Stderr.WriteString("Error writing line\n")
		}

		currPos = nextPos
	}

	result := fmt.Sprintf("Number of unique houses presents delivered at: %d\n", uniqueHouses)
	_, err = buf.WriteString(result)
	if err != nil {
		os.Stderr.WriteString("Error writing final result\n")
	}

	err = buf.Flush()
	if err != nil {
		os.Stderr.WriteString("Error flushing buffer\n")
	}
}

// higher order 32 bits are the X coordinate
// lower order 32 bits are Y coordinate
func makeHouseCoord(x, y int32) houseCoord {
	return houseCoord(uint64(x)<<32 | (uint64(y) & 0xFFFFFFFF))
}

// return only higher order 32 bits as X coordinate
func (h houseCoord) X() int32 {
	return int32(h >> 32)
}

// return only the lower order 32 bits as Y coordinate
func (h houseCoord) Y() int32 {
	temp := h & houseCoord(0xFFFFFFFF)
	return int32(temp)
}
