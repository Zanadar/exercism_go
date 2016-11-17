package gigasecond

import (
	"fmt"
	"time"
)

const testVersion = 4

//AddGigasecond takes a time.Time and add 10^9 seconds to it
func AddGigasecond(t time.Time) time.Time {
	giga := 1000000000
	gigasecond := time.Duration(giga) * time.Second
	fmt.Printf("giga %v, gigasecond %v\n", giga, gigasecond)
	return t.Add(gigasecond)
}
