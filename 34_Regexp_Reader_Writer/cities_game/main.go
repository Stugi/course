package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	cities := []string{
		"Moscow",
		"Washington",
		"New-York",
		"Kiev",
		"Vitebsk",
		"Kishinev",
		"Vladivostok",
	}

	fmt.Println(cities[0])

	for i := 1; i < len(cities); i++ {
		p, n := cities[i-1], cities[i]
		if !isCorrectCityName(p, n) {
			fmt.Println(cities[i], "-", "WRONG!")
			break
		}

		fmt.Println(cities[i], "-", "OK!")
	}
}

var (
	grabLastLetterRe  = regexp.MustCompile(`([a-zA-Z])$`)
	grabFirstLetterRe = regexp.MustCompile(`^([a-zA-Z])`)
)

func isCorrectCityName(prev, next string) bool {
	lastLetterSubmatch := grabLastLetterRe.FindAllStringSubmatch(prev, -1)
	firstLetterSubmatch := grabFirstLetterRe.FindAllStringSubmatch(next, -1)

	if len(lastLetterSubmatch) == 0 || len(firstLetterSubmatch) == 0 {
		return false
	}

	lastLetter := lastLetterSubmatch[0][1]
	firstLetter := firstLetterSubmatch[0][1]

	return lastLetter == strings.ToLower(firstLetter)
}
