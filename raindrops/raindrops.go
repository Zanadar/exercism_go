package raindrops

import (
	"fmt"
	"math"
)

const testVersion = 2

func Convert(n int) string {
	limit := int(math.Sqrt(float64(n)))
	factors := []int{}
	for i := 1; i < limit; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	fmt.Println(factors)
	return "test"
}

// The test program has a benchmark too.  How fast does your Convert convert?
