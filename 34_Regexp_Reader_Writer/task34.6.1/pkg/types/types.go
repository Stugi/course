package mathexample

import (
	"strconv"
	"strings"

	"task34.6.1/pkg/util"
)

type MathExampleFromString struct {
	StartString string
	EndString   string
}

func New(start string) *MathExampleFromString {
	num1, op, num2 := util.ParseMathExpression(start)
	var result int
	num1Int, err := strconv.Atoi(num1)
	if err != nil {
		// handle error
	}
	num2Int, err := strconv.Atoi(num2)
	if err != nil {
		// handle error
	}
	switch op {
	case "+":
		result = num1Int + num2Int
	case "-":
		result = num1Int - num2Int
	}

	return &MathExampleFromString{StartString: start, EndString: strings.Replace(start, "?", strconv.Itoa(result), -1)}
}
