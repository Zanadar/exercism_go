package raindrops

import (
	"math"
	"strconv"
	"strings"
)

const testVersion = 2

//Convert takes in integer and returns a rain patter string based on the factors of the integer.
//
func Convert(n int) string {
	limit := int(math.Sqrt(float64(n)) + 1)
	factors := []int{}

	for i := 1; i < limit; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			if i != n/i {
				factors = append(factors, n/i)
			}
		}
	}

	var rainStrings []string

	for _, factor := range factors {
		switch factor {
		case 3:
			rainStrings = append(rainStrings, "Pling")
		case 5:
			rainStrings = append(rainStrings, "Plang")
		case 7:
			rainStrings = append(rainStrings, "Plong")
		default:
		}

	}
	if len(rainStrings) == 0 {
		return strconv.Itoa(n)
	}
	return strings.Join(rainStrings, "")
}

// The test program has a benchmark too.  How fast does your Convert convert?
