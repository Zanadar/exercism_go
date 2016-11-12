//Package clock gives you some methods for creating, viewing, and modifying "Clock"s.
package clock

import "fmt"
import "time"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

//Clock keeps track of hours and minutes.
type Clock struct {
	Hour   int
	Minute int
}

// New instantiates a Clock with the minutes and seconds you provide.
func New(hour, minute int) *Clock {
	minuteHours := minute / 60
	hour += minuteHours
	minute %= 60
	if hour < 0 {
		hour = 24 + hour%24
	} else {
		hour %= 24
	}
	return &Clock{hour, minute}
}

func (c Clock) String() string {
	clockString := fmt.Sprintf("%v:%v", c.Hour, c.Minute)
	timeObj, err := time.Parse("15:4", clockString)
	if err != nil {
		fmt.Println(err)
	}

	return timeObj.Format("15:04")
}

//Add returns a clock with minutes added
func (c *Clock) Add(minutes int) Clock {
	return Clock{}
}
