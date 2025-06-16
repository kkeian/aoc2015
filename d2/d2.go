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
	var startI uint

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

	var totalSqft uint
	var totalRibbonFeet uint
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
		// calculate ribbon length needed
		switch smallest {
		case lw:
			totalRibbonFeet += 2*present.length + 2*present.width
		case lh:
			totalRibbonFeet += 2*present.length + 2*present.height
		case hw:
			totalRibbonFeet += 2*present.height + 2*present.width
		}
		totalRibbonFeet += calculateRibbonBowLength(present)
	}

	fmt.Printf("The total sqft of wrapping paper needed is: %d sqft\n", totalSqft)
	fmt.Printf("The total ribbon needed is: %d ft\n", totalRibbonFeet)
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
	// configuration for ParseUint
	const base = 10
	const bits = 0 // means int | uint

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

// calculateRibbonBowLength returns the volume of
// a present as the ribbon feet required to tie
// a bow on the present
func calculateRibbonBowLength(r Rectangle) uint {
	return r.height * r.width * r.length
}
