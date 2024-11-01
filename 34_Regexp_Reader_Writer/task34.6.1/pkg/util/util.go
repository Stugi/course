package util

import (
	"regexp"
	"strings"
)

// import "regexp"

// const regex = `^\d+[+-]\d+=\?$`

// // Извлечение примера из текста по регулярному выражению 5+4=?
// func FindExample(text string) (string, bool) {
// 	re, _ := regexp.Compile(regex)

// 	submatch := re.FindString(text)
// 	if len(submatch) == 0 {
// 		return "", false
// 	}

// 	return submatch, true
// }

func ParseMathExpression(expr string) (num1, op, num2 string) {
	// Remove the "=" and "?" characters
	expr = strings.Replace(expr, "=", "", -1)
	expr = strings.Replace(expr, "?", "", -1)

	// Split the expression into parts using a regular expression
	re := regexp.MustCompile(`(\d+)|([-+*/])`)
	parts := re.FindAllString(expr, -1)

	// Extract the numbers and operator
	num1 = parts[0]
	op = parts[1]
	num2 = parts[2]

	return
}
