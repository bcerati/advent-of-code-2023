package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	puzzle, error := ioutil.ReadFile("puzzle")
	if error != nil {
		panic("Cant read puzzle file")
	}

	result := 0
	lines := strings.Split(string(puzzle), "\n")

	stringNumber := ""
	for idxLine, line := range lines {
		for idxCol, char := range line {
			if unicode.IsDigit(char) {
				stringNumber += string(char)
			} else if stringNumber != "" {
				if isPartOf(lines, stringNumber, idxLine, idxCol) {
					n, _ := strconv.Atoi(stringNumber)
					result += n
				}

				stringNumber = ""
			}
		}
	}

	fmt.Println(result)
}

func charToInt(c rune) int {
	n, _ := strconv.Atoi(string(c))

	return n
}

func isPartOf(lines []string, number string, idxLine int, idxCol int) bool {
	fromIdx := idxCol - len(number) - 1
	toIdx := idxCol

	if fromIdx < 0 {
		fromIdx = 0
	}

	if toIdx >= len(lines[0]) {
		toIdx = len(lines[0]) - 1
	}

	if idxCol == 0 {
		idxLine = idxLine - 1
		toIdx = len(lines[0]) - 1
		fromIdx = toIdx - len(number) - 1
	}

	// checking line -1
	if idxLine-1 > -1 {
		if hasSymbol(lines[idxLine-1][fromIdx : toIdx+1]) {
			return true
		}
	}

	// checking line + 1
	if len(lines[idxLine+1]) > 0 {
		if hasSymbol(lines[idxLine+1][fromIdx : toIdx+1]) {
			return true
		}
	}

	if hasSymbol(string(lines[idxLine][fromIdx : toIdx+1])) {
		return true
	}

	return false
}

func hasSymbol(str string) bool {
	for _, c := range str {
		if unicode.IsDigit(c) || string(c) == "." {
			continue
		}

		return true
	}

	return false
}
