package is

import (
	"fmt"
	"github.com/corbym/gocrest"
	"strings"
)

// StringContaining finds if all x's are contained as value in y.
// Acts like "ContainsAll", all elements given must be present.
func StringContaining(expected ...string) *gocrest.Matcher[string] {
	match := new(gocrest.Matcher[string])
	match.Describe = fmt.Sprintf("something that contains %v", expected)
	match.Matches = func(actual string) bool {
		for _, e := range expected {
			if !strings.Contains(actual, e) {
				return false
			}
		}
		return true
	}
	return match
}

// MapContaining finds if all map[k] 's value V is contained as a value of actual[k]
// Acts like "ContainsAll", all elements given must be present in actual in the same order as the expected values.
func MapContaining[K comparable, V comparable](expected map[K]V) *gocrest.Matcher[map[K]V] {
	match := new(gocrest.Matcher[map[K]V])
	match.Describe = fmt.Sprintf("something that contains %v", expected)
	match.Matches = func(actual map[K]V) bool {
		return mapActualContainsExpected(expected, actual)
	}
	return match
}

// MapContainingValues finds if all values V is contained as a value of actual[k]
// Acts like "ContainsAll", all elements given must be present in actual in the same order as the expected values.
func MapContainingValues[K comparable, V comparable](expected ...V) *gocrest.Matcher[map[K]V] {
	match := new(gocrest.Matcher[map[K]V])
	match.Describe = fmt.Sprintf("something that contains %v", expected)
	match.Matches = func(actual map[K]V) bool {
		return mapActualContainsExpectedValues(expected, actual)
	}
	return match
}

// MapMatchingValues finds if all values V is match a value of actual[k]
// Acts like "ContainsAll", all elements given must match in actual in the same order as the expected values.
func MapMatchingValues[K comparable, V comparable](expected ...*gocrest.Matcher[V]) *gocrest.Matcher[map[K]V] {
	match := new(gocrest.Matcher[map[K]V])
	match.Describe = descriptionForMatchers(expected)
	match.Matches = func(actual map[K]V) bool {
		return mapActualMatchesExpected(expected, actual)
	}
	return match
}

func descriptionForMatchers[A any](expected []*gocrest.Matcher[A]) string {
	var description = ""
	for x, m := range expected {
		description += m.Describe
		if x < len(expected)-1 {
			description += " and "
		}
	}
	return description
}

// ArrayContaining finds if all x's are contained in y.
// Acts like "ContainsAll", all elements given must be present in actual.
func ArrayContaining[A comparable](expected ...A) *gocrest.Matcher[[]A] {
	match := new(gocrest.Matcher[[]A])
	match.Describe = fmt.Sprintf("something that contains %v", descriptionFor(expected))
	match.Matches = func(actual []A) bool {
		return listContains(expected, actual)
	}
	return match
}

// ArrayMatching finds if all x's are matched in y.
// Acts like "ContainsAll", all elements given must be present in actual.
func ArrayMatching[A comparable](expected ...*gocrest.Matcher[A]) *gocrest.Matcher[[]A] {
	match := new(gocrest.Matcher[[]A])
	match.Describe = fmt.Sprintf("something that contains %v", descriptionFor(expected))
	match.Matches = func(actual []A) bool {
		return listMatches(expected, actual)
	}
	return match
}
func mapActualContainsExpected[K comparable, V comparable](expected map[K]V, actual map[K]V) bool {
	for k, ev := range expected {
		if actual[k] != ev {
			return false
		}
	}
	return true
}
func mapActualContainsExpectedValues[K comparable, V comparable](expected []V, actual map[K]V) bool {
	// Build a multiset of available actual values so duplicate expected values require duplicate actuals.
	available := make(map[V]int)
	for _, v := range actual {
		available[v]++
	}
	for _, e := range expected {
		if available[e] == 0 {
			return false
		}
		available[e]--
	}
	return true
}
func mapActualMatchesExpected[K comparable, V comparable](expected []*gocrest.Matcher[V], actual map[K]V) bool {
	// Greedy: each expected matcher consumes one unique actual key so duplicates are handled correctly.
	used := make(map[K]bool)
	for _, exp := range expected {
		found := false
		for k, v := range actual {
			if !used[k] && exp.Matches(v) {
				used[k] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func listContains[T comparable, A []T](expected A, actualValue A) bool {
	// Multiset matching: duplicate expected elements require the same count in actual.
	available := make(map[T]int)
	for _, act := range actualValue {
		available[act]++
	}
	for _, exp := range expected {
		if available[exp] == 0 {
			return false
		}
		available[exp]--
	}
	return true
}
func listMatches[T comparable](expected []*gocrest.Matcher[T], actualValue []T) bool {
	// Greedy: each expected matcher consumes one unmatched actual element.
	used := make([]bool, len(actualValue))
	for _, exp := range expected {
		found := false
		for i, act := range actualValue {
			if !used[i] && exp.Matches(act) {
				used[i] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func descriptionFor[T any, A []T](expected A) string {
	var description = ""
	for x, e := range expected {
		description += fmt.Sprintf("<%v>", e)
		if x < len(expected)-1 {
			description += " and "
		}
	}
	return description
}
