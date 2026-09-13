package is

import (
	"fmt"
	"github.com/corbym/gocrest"
)

// SameInstance returns a matcher that passes when actual and expected are the same
// instance (pointer identity via ==). For non-pointer comparable types this behaves
// identically to EqualTo; its primary use is asserting two variables reference the
// same object in memory.
func SameInstance[A comparable](expected A) *gocrest.Matcher[A] {
	match := new(gocrest.Matcher[A])
	match.Describe = fmt.Sprintf("the same instance as <%v>", expected)
	match.Matches = func(actual A) bool {
		match.Actual = ""
		if actual == expected {
			return true
		}
		match.Actual = fmt.Sprintf("%v", actual)
		return false
	}
	return match
}
