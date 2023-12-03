package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	result := 0
	puzzle, error := ioutil.ReadFile("puzzle")
	if error != nil {
		panic("Cant read puzzle file")
	}

	lines := strings.Split(string(puzzle), "\n")

	regex := regexp.MustCompile("Game ([0-9]+):")
	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)

		if len(matches) == 0 {
			continue
		}

		regexRed := regexp.MustCompile("([0-9]+) red")
		matchesRed := regexRed.FindAllStringSubmatch(line, -1)

		regexGreen := regexp.MustCompile("([0-9]+) green")
		matchesGreen := regexGreen.FindAllStringSubmatch(line, -1)

		regexBlue := regexp.MustCompile("([0-9]+) blue")
		matchesBlue := regexBlue.FindAllStringSubmatch(line, -1)

		maxRed := getMax(matchesRed)
		maxGreen := getMax(matchesGreen)
		maxBlue := getMax(matchesBlue)

		result += maxRed * maxGreen * maxBlue
	}

	fmt.Println(result)
}

func getMax(matches [][]string) int {
	max := 1
	for _, game := range matches {
		n, _ := strconv.Atoi(game[1])

		if n > max {
			max = n
		}
	}
	return max
}
