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
	result2 := 0
	var gears map[string][]string
	gears = make(map[string][]string)

	lines := strings.Split(string(puzzle), "\n")

	stringNumber := ""
	for idxLine, line := range lines {
		for idxCol, char := range line {
			if unicode.IsDigit(char) {
				stringNumber += string(char)
			} else if stringNumber != "" {
				if isPartOf(lines, stringNumber, idxLine, idxCol, &gears) {
					n, _ := strconv.Atoi(stringNumber)
					result += n
				}

				stringNumber = ""
			}
		}
	}

	for _, numbers := range gears {
		if len(numbers) == 2 {
			n1, _ := strconv.Atoi(numbers[0])
			n2, _ := strconv.Atoi(numbers[1])

			result2 += n1 * n2
		}
	}
	fmt.Println("Part1 = ", result)
	fmt.Println("Part2 = ", result2)
}

func charToInt(c rune) int {
	n, _ := strconv.Atoi(string(c))

	return n
}

func isPartOf(lines []string, number string, idxLine int, idxCol int, gears *map[string][]string) bool {
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

	symbols := ""
	// checking line -1
	if idxLine-1 > -1 {
		symbols = lines[idxLine-1][fromIdx : toIdx+1]
		if hasSymbol(symbols) {
			appendForSymbols(number, symbols, idxLine-1, fromIdx, gears)
			return true
		}
	}

	// checking line + 1
	if len(lines[idxLine+1]) > 0 {
		symbols = lines[idxLine+1][fromIdx : toIdx+1]
		if hasSymbol(symbols) {
			appendForSymbols(number, symbols, idxLine+1, fromIdx, gears)
			return true
		}
	}

	symbols = string(lines[idxLine][fromIdx : toIdx+1])
	if hasSymbol(string(lines[idxLine][fromIdx : toIdx+1])) {
		appendForSymbols(number, symbols, idxLine, fromIdx, gears)
		return true
	}

	return false
}

func appendForSymbols(number string, symbols string, idxLine int, idxCol int, gears *map[string][]string) {
	idxLineStr := strconv.Itoa(idxLine)

	for k, c := range symbols {
		if string(c) == "*" {
			idxColStr := strconv.Itoa(idxCol + k)

			(*gears)[idxLineStr+";"+idxColStr] = append((*gears)[idxLineStr+";"+idxColStr], number)
		}
	}
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
