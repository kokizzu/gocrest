# 009 — `is.StringContainingInOrder`

**Status:** backlog  
**Package:** `is`  
**Hamcrest parity:** yes (`stringContainsInOrder`)

## Summary

Add `is.StringContainingInOrder(substrings...)` that passes when all substrings appear in the string left-to-right in the given order.

The existing `is.StringContaining` checks all substrings are present but ignores order.
This matcher adds the ordering constraint.

## API

```go
func StringContainingInOrder(expected ...string) *gocrest.Matcher[string]
```

## Example usage

```go
then.AssertThat(t, "foo bar baz", is.StringContainingInOrder("foo", "bar", "baz"))
then.AssertThat(t, "one two three", is.StringContainingInOrder("one", "three"))
```

## Acceptance criteria

- Passes when all substrings appear and each starts after the end of the previous match
- Fails when any substring is absent
- Fails when substrings appear out of order
- Single substring behaves identically to `is.StringContaining`
- Failure description: `a string containing ["foo" "bar" "baz"] in order`

## Test shape

```go
func TestStringContainingInOrder(t *testing.T) {
    tests := []struct {
        actual     string
        expected   []string
        shouldFail bool
    }{
        {actual: "foo bar baz",  expected: []string{"foo", "bar", "baz"}, shouldFail: false},
        {actual: "foo bar baz",  expected: []string{"foo", "baz", "bar"}, shouldFail: true},
        {actual: "foo bar baz",  expected: []string{"foo", "qux"},        shouldFail: true},
        {actual: "abcabc",       expected: []string{"abc", "abc"},        shouldFail: false},
        {actual: "hello world",  expected: []string{"hello"},             shouldFail: false},
    }
    for _, test := range tests {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, test.actual, is.StringContainingInOrder(test.expected...))
        then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
    }
}
```
