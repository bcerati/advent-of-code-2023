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

		gameNumber, _ := strconv.Atoi(matches[1])

		regexRed := regexp.MustCompile("([0-9]+) red")
		matchesRed := regexRed.FindAllStringSubmatch(line, -1)

		regexGreen := regexp.MustCompile("([0-9]+) green")
		matchesGreen := regexGreen.FindAllStringSubmatch(line, -1)

		regexBlue := regexp.MustCompile("([0-9]+) blue")
		matchesBlue := regexBlue.FindAllStringSubmatch(line, -1)

		addToResult := gameNumber
		for _, set := range matchesRed {
			if asMoreThanItShould(set, 12) {
				addToResult = 0
				break
			}
		}

		for _, set := range matchesGreen {
			if asMoreThanItShould(set, 13) {
				addToResult = 0
				break
			}
		}

		for _, set := range matchesBlue {
			if asMoreThanItShould(set, 14) {
				addToResult = 0
				break
			}
		}

		result += addToResult
	}

	fmt.Println(result)
}

func asMoreThanItShould(matches []string, n int) bool {
	if len(matches) > 1 {
		v, _ := strconv.Atoi(matches[1])
		return !(v <= n)
	}

	panic("This should not happen!")
}
