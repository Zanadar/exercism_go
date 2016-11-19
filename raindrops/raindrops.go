package raindrops

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const testVersion = 2

func Convert(n int) string {
	fmt.Println("************** new Call to convert n is: ", n)
	limit := int(math.Sqrt(float64(n)) + 1)
	factors := []int{}

	for i := 1; i < limit; i++ {
		if n%i == 0 {
			fmt.Println("Factor added: ", i)
			factors = append(factors, i)
			if i != n/i {
				fmt.Println("Factor added: ", n/i)
				factors = append(factors, n/i)
			}
		}
	}

	var rainStrings []string

	for i, factor := range factors {
		fmt.Println(i, factor, factors)
		switch factor {
		case 3:
			rainStrings = append(rainStrings, "Pling")
			fmt.Println("3 factor", rainStrings)
		case 5:
			rainStrings = append(rainStrings, "Plang")
			fmt.Println("5 factor", rainStrings)
		case 7:
			rainStrings = append(rainStrings, "Plong")
			fmt.Println("7 factor", rainStrings)
		default:
			fmt.Println("Garbage")
		}

	}
	if len(rainStrings) == 0 {
		return strconv.Itoa(n)
	}
	fmt.Println(n, rainStrings)
	return strings.Join(rainStrings, "")
}

// The test program has a benchmark too.  How fast does your Convert convert?
