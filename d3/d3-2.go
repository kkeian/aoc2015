package main

import (
	"fmt"
	"os"
)

type coordinate struct {
	X, Y int
}

type move int

const (
	Up move = iota
	Down
	Left
	Right
)

var direction = map[byte]move{
	'^': Up,
	'v': Down,
	'>': Right,
	'<': Left,
}

type trackingPoint map[coordinate]bool

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		os.Stderr.WriteString("Error reading input")
		os.Exit(1)
	}

	startingPoint := coordinate{0, 0}
	santaLoc, roboSantaLoc := startingPoint, startingPoint
	santasTurn := true
	deliveryMan := map[bool]*coordinate{
		true:  &santaLoc,
		false: &roboSantaLoc,
	}

	visitedHouses := trackingPoint{startingPoint: true}

	uniqueHouses := 1
	for _, symbol := range content {
		deliverer := deliveryMan[santasTurn]
		switch direction[symbol] {
		case Up:
			deliverer.Y++
		case Down:
			deliverer.Y--
		case Right:
			deliverer.X++
		case Left:
			deliverer.X--
		}

		if !visitedHouses[*deliverer] {
			visitedHouses[*deliverer] = true
			uniqueHouses++
		}

		santasTurn = !santasTurn
	}
	fmt.Printf("%d total unique houses had presents delivered\n", uniqueHouses)
}
