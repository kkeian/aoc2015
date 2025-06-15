package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err.Error())
	}

	boxes := []Rectangle{}
	var startI uint = 0

	for range content {
		box := new(Rectangle)
		offset, err := parseDimensions(content[startI:], box)
		if err != nil {
			fmt.Printf("Error reading file.\n")
			os.Exit(1)
		}
		startI += offset
		boxes = append(boxes, *box)
	}

	var totalSqft uint = 0
	for _, present := range boxes {
		lw := present.length * present.width
		lh := present.length * present.height
		hw := present.height * present.width
		totalSqft += 2 * (lw + lh + hw)
		smallest := lw
		if lh < smallest {
			smallest = lh
		}
		if hw < smallest {
			smallest = hw
		}
		totalSqft += smallest
	}

	fmt.Printf("The total sqft of wrapping paper needed is: %d sqft\n", totalSqft)
}

type Rectangle struct {
	length, width, height uint
}

func (r Rectangle) String() string {
	return fmt.Sprintf("%dx%dx%d", r.length, r.width, r.height)
}

func parseDimensions(bytes []byte, box *Rectangle) (nextStartIndex uint, err error) {
	// parse box dimensions
	start, xs := 0, 0
	base := 10
	bits := 0 // means int | uint

	for i, sym := range bytes {
		switch sym {
		case '\n':
			temp, err := strconv.ParseUint(string(bytes[start:i]), base, bits)
			if err != nil {
				return uint(i), err
			}
			dim := uint(temp)
			box.height = dim
			return uint(i + 1), nil
		case 'x', 'X':
			temp, err := strconv.ParseUint(string(bytes[start:i]), base, bits)
			if err != nil {
				return uint(i), err
			}
			dim := uint(temp)
			switch xs {
			case 0:
				box.length = dim
			case 1:
				box.width = dim
			}
			start = i + 1
			xs += 1
		}
	}
	return 0, nil
}
