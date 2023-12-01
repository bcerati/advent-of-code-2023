package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	puzzle, error := ioutil.ReadFile("puzzle")
	if error != nil {
		panic("Cant read puzzle file")
	}

	lines := strings.Split(string(puzzle), "\n")
	result := 0

	for _, line := range lines {
		currentCalibration := []string{}

		for charIndex, char := range line {
			_, err := strconv.Atoi(string(char))

			// case of an integer
			if err == nil {
				currentCalibration = append(currentCalibration, string(char))
			} else {
				for n := range numbers {
					// managing out of bounds
					if charIndex+len(n)-1 < len(line) {
						i, exists := numbers[line[charIndex:charIndex+len(n)]]
						if exists {
							currentCalibration = append(currentCalibration, strconv.Itoa(i))
						}
					}
				}
			}
		}

		if len(currentCalibration) == 0 {
			continue
		}

		first := currentCalibration[0]
		second := first

		if len(currentCalibration) > 1 {
			second = currentCalibration[len(currentCalibration)-1]
		}

		conv, _ := strconv.Atoi(first + second)
		result += conv
	}

	fmt.Println("Result = ", result)
	// 56042
	// 55358
}
