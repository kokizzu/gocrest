package has

import (
	"fmt"
	"github.com/corbym/gocrest"
)

// EveryElement Checks whether the nth element of the array/slice matches the nth expectation passed
func EveryElement[A any](expects ...*gocrest.Matcher[A]) *gocrest.Matcher[[]A] {
	match := new(gocrest.Matcher[[]A])
	match.Describe = fmt.Sprintf("elements to match %s", describe(expects, "and"))

	match.Matches = func(actual []A) bool {
		match.Actual = "" // reset for this invocation
		if len(actual) != len(expects) {
			return false
		}
		result := true
		for i, act := range actual {
			if !expects[i].Matches(act) {
				result = false
			}
			match.AppendActual(expects[i].Actual)
		}
		return result
	}

	return match
}

func describe[A any](matchers []*gocrest.Matcher[A], conjunction string) string {
	var description string
	for x := 0; x < len(matchers); x++ {
		description += fmt.Sprintf("[%v]:%v", x, matchers[x].Describe)
		if x+1 < len(matchers) {
			description += fmt.Sprintf(" %s ", conjunction)
		}
	}
	return description
}
