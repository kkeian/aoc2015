package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	floorNum := 0
	var res int

	for i, sym := range content {
		if sym == '(' {
			floorNum += 1
		} else if sym == ')' {
			floorNum -= 1
		} else {
			fmt.Println("This should not be happening BLAAAAARGH!")
		}

		if floorNum == -1 {
			res = i + 1
			break
		}
	}

	fmt.Printf("The first symbol that puts Santa on floor -1 (basement) is: %d\n", res)
}
