package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	puzzle, error := ioutil.ReadFile("puzzle")
	if error != nil {
			panic("Cant read puzzle file")
	}

	results := 0
	scratches := map[string][]string{}

	lines := strings.Split(string(puzzle), "\n")

	cardRegex := regexp.MustCompile(`^[A-Za-z: ]+([0-9]+): ([0-9 ]+) \| ([0-9 ]+)$`)
	for _, line := range lines {
		matchesCard := cardRegex.FindStringSubmatch(line)

		cardNumber := matchesCard[1]
		winningNumbers := strings.Split(matchesCard[2], " ")
		hand := strings.Split(matchesCard[3], " ")

		intersect := sliceIntersect(winningNumbers, hand)

		if len(intersect) > 0 {
			results += int(math.Pow(2, float64(len(intersect) - 1)))
			
			scratches[cardNumber] = process(cardNumber, intersect)
		} else {
			scratches[cardNumber] = []string{}
		}
	}

	allScratches := []string{}
	copies := []string{}
	for k, s := range scratches {
		allScratches = append(allScratches, k)
		for _, c := range s {
			copies = append(copies, c)
		}
	}

	for len(copies) > 0 {
		for k, n := range copies {
			allScratches = append(allScratches, n)
			copies = removeIndex(k, copies)

			values, ok := scratches[n]

			if ok {
				for _, v := range values {
					copies = append(copies, v)
				} 
			}
			break;
		}
	}
	fmt.Println("allScratches = ", len(allScratches))
}

func sliceIntersect(slice1 []string, slice2 []string) []string {
	results := []string{}

	for _, s1 := range slice1 {
		for _, s2 := range slice2 {
			if (s1 == s2 && s1 != "" && s2 != "") {
				results = append(results, s1)
			}
		}
	}

	return results;
}

func process(cardNumber string, cards []string)[]string {
	n, _ := strconv.Atoi(cardNumber)

	scratches := []string{}
	for k := range cards {
		scratches = append(scratches, fmt.Sprint(n + k + 1))
	}

	return scratches
}

func removeIndex(i int, s []string) []string {
	s[i] = s[len(s)-1]
	s[len(s)-1] = ""
	s = s[:len(s)-1]
	
	return s
}