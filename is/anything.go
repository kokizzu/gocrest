package is

import "github.com/corbym/gocrest"

// Anything returns a matcher that always passes regardless of the actual value.
// Useful as a placeholder in composite matchers where one position is a "don't care".
func Anything[A any]() *gocrest.Matcher[A] {
	return &gocrest.Matcher[A]{
		Describe: "anything",
		Matches:  func(actual A) bool { return true },
	}
}
