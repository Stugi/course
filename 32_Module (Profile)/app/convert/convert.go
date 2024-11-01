package convert

import "math"

func MetersToFeet(meters int) (feet int) {
	feet = meters * int(math.Round(3.28084))

	return
}
