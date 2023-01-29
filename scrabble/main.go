
package main

import (
	"fmt"
	"unicode"
)

type PointCalculation struct {
	points int
	str    string
}

var pMap = map[rune]int{
	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
	'D': 2, 'G': 2,
	'B': 3, 'C': 3, 'M': 3, 'P': 3,
	'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
	'K': 5,
	'J': 8, 'X': 8,
	'Q': 10, 'Z': 10,
}

func Score(val string) int {
	var sum int
	if len(val) == 0 {
		return 0
	}
	for _, char := range val {
		upperChar := unicode.ToUpper(char)
		if points, ok := pMap[upperChar]; ok {
			sum += points
		}
	}
	return sum
}

func main() {
	fmt.Println("Scrabble Project")
	fmt.Println("Calculate", Score(""))
	fmt.Println("Calculate", Score("abcdefghijklmnopqrstuvwxyz"))
}

