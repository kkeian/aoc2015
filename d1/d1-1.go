package main

import (
	"fmt"
	"os"
)

func main() {
	contents, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	floorNum := 0
	
	for _, symbol := range contents {
		if symbol == '(' {
			floorNum += 1
			// fmt.Println("UP")
		} else if symbol == ')' {
			floorNum -= 1
			// fmt.Println("DOWN")
		} else {
			fmt.Println("This should not be happening.")
		}
	}

	fmt.Printf("Santa ends up at floor: %d\n", floorNum)
}
