package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	dates := []string{
		"12.09.1978",
		"1990/06/10",
		"08.03.2021",
		"12.04.1986",
		"25 dec 1988",
		"2001/05/25",
	}

	for _, d := range dates {
		year, month, day, err := parseDate(d)
		if err != nil {
			fmt.Println("ERROR!", err, "-", d)
			continue
		}

		fmt.Printf("%04d.%02d.%02d\n", year, month, day)
	}
}

var UndefinedDateFormat = errors.New("undefined date format")

func parseDate(date string) (year, month, day int64, err error) {
	// TODO: try dd.mm.YYYY format

	// TODO: or try YYYY/mm/dd format

	// or error
	err = UndefinedDateFormat

	return
}

func strToInt(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}
