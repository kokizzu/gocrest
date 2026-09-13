package is

import (
	"fmt"
	"github.com/corbym/gocrest"
	"reflect"
)

// Zero returns a matcher that passes when actual is the zero value for its type:
// 0 for numeric types, "" for strings, false for booleans, nil for pointers, etc.
func Zero[A any]() *gocrest.Matcher[A] {
	match := new(gocrest.Matcher[A])
	match.Describe = "zero value"
	match.Matches = func(actual A) bool {
		match.Actual = ""
		// reflect.ValueOf(&actual).Elem() preserves the interface wrapper so that
		// a typed-nil interface (e.g. (*MyError)(nil) stored as error) is not
		// treated as zero — consistent with Go's own == nil semantics.
		v := reflect.ValueOf(&actual).Elem()
		if !v.IsValid() || v.IsZero() {
			return true
		}
		match.Actual = fmt.Sprintf("%v", actual)
		return false
	}
	return match
}
