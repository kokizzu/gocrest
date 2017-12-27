package gocrest

import "fmt"

//Takes some matchers and checks if all the matchers return true
//returns a matcher that performs the the test on the input matchers
func AllOf(allMatchers ... *Matcher) (*Matcher) {
	matcher := new(Matcher)
	matcher.matches = func(actual interface{}) bool {
		matcher.describe = fmt.Sprintf("all of (%s)", describe(allMatchers, "and"))
		for x := 0; x < len(allMatchers); x++ {
			if !allMatchers[x].matches(actual) {
				return false
			}
		}
		return true
	}
	return matcher
}
