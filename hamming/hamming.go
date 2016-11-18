package hamming

import "errors"

const testVersion = 5

//Distance returns how many characters are different from their equivalent in the other string.
//It expects strings of equal length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Strings must be of equal length.")
	}
	var diff int
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff, nil
}
