package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		os.Stderr.WriteString("Error reading file")
		os.Exit(1)
	}

	hash := md5.New()
	// Desired 0x00000X - X is don't care
	// 		any number < 0x00001F is 0x00000F or lower which == 0x00000X
	// targetPrefix := []byte{0, 0, 0x1F} // part 1
	targetPrefix := []byte{0, 0, 0} // part 2
	var hashInput []byte
	// smallest numerical suffix with first 5 digits in checksum == 0
	var suffix uint64
	for suffix = uint64(1); ; suffix++ {
		hash.Reset()
		hashInput = bytes.Clone(content)
		byteStr := U64ToBytes(suffix)
		hashInput = append(hashInput, byteStr...)
		_, err := hash.Write(hashInput)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error writing hash")
			os.Exit(1)
		}
		checksumPrefix := hash.Sum(nil)[:3]
		done := bytes.Compare(checksumPrefix, targetPrefix)
		if done < 1 {
			break
		}
	}

	fmt.Printf("md5 hash of %s: %x\n", hashInput, hash.Sum(nil))
	fmt.Printf("The smallest number suffix is: %d\n", suffix)
}

// For fast conversion of digit to ascii character
// equivalent
var numberTable = map[uint64]byte{
	0: byte('0'),
	1: byte('1'),
	2: byte('2'),
	3: byte('3'),
	4: byte('4'),
	5: byte('5'),
	6: byte('6'),
	7: byte('7'),
	8: byte('8'),
	9: byte('9'),
}

// Helper for appending numerical suffix to
// input byte slice. MD5 calculated on byte slice
// which is UTF-8 string data.
//
// Converts num into its its byte slice
// form where each element is the ascii character
// for the digit
func U64ToBytes(num uint64) []byte {
	var divisor uint64 = 10
	var result []byte
	for num/divisor > uint64(0) {
		dig := num % divisor
		result = append([]byte{numberTable[dig]}, result...)
		num /= divisor
	}
	// make sure 1 digit num and final digit of num
	// are pre-pended
	result = append([]byte{numberTable[num]}, result...)
	return result
}
