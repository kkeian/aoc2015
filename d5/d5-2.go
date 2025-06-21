package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		os.Stderr.WriteString("Error reading file")
		os.Exit(1)
	}

	matches := 0
	twoPair, charSandwich := false, false
	twoPairEOL := false
	// debug only:
	var matchingLines []int
	for i := 0; i < len(content)-2; i++ {
		if !twoPairEOL && content[i+3] == '\n' {
			twoPairEOL = true
		}
		if !twoPair && !twoPairEOL {
			// scan for two pair non-overlapping
			// O(n^2) - not sure if two pair find can be made faster
			//		  without regexp module
			for j := i + 2; content[j+1] != '\n'; j++ {
				if slices.Equal(content[i:i+2], content[j:j+2]) {
					twoPair = true
					break
				}
			}
		}
		// char sandwich e.g. cac
		if content[i+2] != '\n' {
			if !charSandwich && (content[i] == content[i+2]) {
				charSandwich = true
			}
			continue // where this is located is important
		}

		if twoPair && charSandwich {
			matches++
			// debug only:
			matchingLines = append(matchingLines, (i/17)+1)
		}
		twoPair, charSandwich = false, false
		twoPairEOL = false
		// jump to start of next line
		i = i + 2 // content[i] == '\n' b/c of loop incrementer
	}

	// debug only:
	for _, line := range matchingLines {
		fmt.Println(line)
	}

	fmt.Printf("%d matching lines\n", matches)
}
