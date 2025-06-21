package main

import (
	"fmt"
	"os"
)

var vowels = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

var badStrings = map[[2]byte]bool{
	{'a', 'b'}: true,
	{'c', 'd'}: true,
	{'p', 'q'}: true,
	{'x', 'y'}: true,
}

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		os.Stderr.WriteString("Error reading input file\n")
		os.Exit(1)
	}

	totalNice := 0
	vowelCount := 0
	oneDupChar := false
	var prevChar byte
	skipLine := false
	for _, char := range content {
		if char != '\n' {
			if skipLine {
				continue
			}
			if badStrings[[2]byte{prevChar, char}] {
				skipLine = true
				continue
			}
			if vowels[char] {
				vowelCount++
			}
			if !oneDupChar && prevChar == char {
				oneDupChar = true
			}
			prevChar = char
			continue
		}

		if !skipLine && (vowelCount > 2 && oneDupChar) {
			totalNice++
		}
		vowelCount = 0
		oneDupChar = false
		prevChar = 0
		skipLine = false
	}

	fmt.Printf("Found %d nice strings\n", totalNice)
}
