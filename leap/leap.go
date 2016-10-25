package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// IsLeapYear tests if a number is a leap year.
// A leap year is divisible by 4, except if divisible by 100 unless also divisible by 400
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
