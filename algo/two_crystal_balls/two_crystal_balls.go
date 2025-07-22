package two_crystal_balls

import (
	"math"
)

// twoCrystalBalls returns the first index where true appears, or -1 if not found.
func twoCrystalBalls(breaks []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jumpAmount

	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	start := i - jumpAmount

	for j := start; j <= i && j < len(breaks); j++ {
		if breaks[j] {
			return j
		}
	}

	return -1
}
