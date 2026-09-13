# 008 — `is.BlankString`

**Status:** backlog  
**Package:** `is`  
**Hamcrest parity:** yes (`blankString`, `blankOrNullString`)

## Summary

Add `is.BlankString()` that passes when a string is empty or contains only whitespace (`strings.TrimSpace(actual) == ""`).
This is distinct from `is.EmptyString()` (which requires `len == 0`) and fills a Hamcrest gap.

## API

```go
func BlankString() *gocrest.Matcher[string]
```

## Example usage

```go
then.AssertThat(t, "",    is.BlankString())
then.AssertThat(t, "   ", is.BlankString())
then.AssertThat(t, "\t\n", is.BlankString())
```

## Acceptance criteria

- Passes for `""`
- Passes for strings containing only spaces, tabs, newlines
- Fails for strings with any non-whitespace character
- Failure description: `a blank string`
- Actual in failure: `<actual>`

## Test shape

```go
func TestBlankString(t *testing.T) {
    tests := []struct {
        actual     string
        shouldFail bool
    }{
        {actual: "",      shouldFail: false},
        {actual: "   ",   shouldFail: false},
        {actual: "\t",    shouldFail: false},
        {actual: "\n",    shouldFail: false},
        {actual: " a ",   shouldFail: true},
        {actual: "hello", shouldFail: true},
    }
    for _, test := range tests {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, test.actual, is.BlankString())
        then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
    }
}
```
