package hello

import (
	"strings"
)

const testVersion = 2

//HelloWorld is the main function
func HelloWorld(name string) string {
	if name == "" {
		name += "World"
	}
	parts := []string{"Hello, ", name, "!"}
	return strings.Join(parts, "")
}
