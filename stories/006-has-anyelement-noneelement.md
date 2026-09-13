# 006 — `has.AnyElement` / `has.NoneElement`

**Status:** backlog  
**Package:** `has`  
**Hamcrest parity:** yes (`hasItem` maps to `AnyElement`; `not(hasItem(...))` maps to `NoneElement`)

## Summary

Logical complements of the existing `has.EveryElement`:

- `has.AnyElement(matcher)` — passes if at least one element satisfies the matcher (Hamcrest: `hasItem`)
- `has.NoneElement(matcher)` — passes if no element satisfies the matcher

## API

```go
func AnyElement[A any](expected *gocrest.Matcher[A]) *gocrest.Matcher[[]A]
func NoneElement[A any](expected *gocrest.Matcher[A]) *gocrest.Matcher[[]A]
```

## Example usage

```go
then.AssertThat(t, []int{1, 2, 3}, has.AnyElement(is.GreaterThan(2)))
then.AssertThat(t, []string{"foo", "bar"}, has.NoneElement(is.EqualTo("baz")))
```

## Acceptance criteria

### `has.AnyElement`
- Passes when at least one element matches
- Fails when no element matches
- Fails on empty slice
- Failure description: `a slice with any element matching <matcher.Describe>`

### `has.NoneElement`
- Passes when no element matches
- Passes on empty slice
- Fails when any element matches
- Failure description: `a slice with no element matching <matcher.Describe>`

## Test shape

```go
func TestAnyElement(t *testing.T) {
    tests := []struct {
        actual     []int
        matcher    *gocrest.Matcher[int]
        shouldFail bool
    }{
        {actual: []int{1, 2, 3}, matcher: is.GreaterThan(2),  shouldFail: false},
        {actual: []int{1, 2, 3}, matcher: is.GreaterThan(10), shouldFail: true},
        {actual: []int{},        matcher: is.EqualTo(1),      shouldFail: true},
    }
    for _, test := range tests {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, test.actual, has.AnyElement(test.matcher))
        then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
    }
}

func TestNoneElement(t *testing.T) {
    tests := []struct {
        actual     []string
        matcher    *gocrest.Matcher[string]
        shouldFail bool
    }{
        {actual: []string{"foo", "bar"}, matcher: is.EqualTo("baz"), shouldFail: false},
        {actual: []string{},             matcher: is.EqualTo("foo"), shouldFail: false},
        {actual: []string{"foo", "bar"}, matcher: is.EqualTo("foo"), shouldFail: true},
    }
    for _, test := range tests {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, test.actual, has.NoneElement(test.matcher))
        then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
    }
}
```
