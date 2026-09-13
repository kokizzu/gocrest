# 004 — `is.EqualToIgnoringCase`

**Status:** backlog  
**Package:** `is`  
**Hamcrest parity:** yes (`equalToIgnoringCase`)

## Summary

Add `is.EqualToIgnoringCase(expected)` as a natural sibling to the existing `is.EqualToIgnoringWhitespace`.
Uses `strings.EqualFold` for Unicode-correct case folding.

## API

```go
func EqualToIgnoringCase(expected string) *gocrest.Matcher[string]
```

## Example usage

```go
then.AssertThat(t, "Hello World", is.EqualToIgnoringCase("hello world"))
then.AssertThat(t, "GOCREST", is.EqualToIgnoringCase("gocrest"))
```

## Acceptance criteria

- Passes when strings are equal ignoring case (`strings.EqualFold`)
- Fails when strings differ beyond case
- Failure description: `ignoring case value equal to <expected>`
- Actual in failure: `<actual>`

## Test shape

```go
func TestEqualToIgnoringCase(t *testing.T) {
    tests := []struct {
        actual     string
        expected   string
        shouldFail bool
    }{
        {actual: "Hello", expected: "hello",   shouldFail: false},
        {actual: "HELLO", expected: "hello",   shouldFail: false},
        {actual: "hello", expected: "hello",   shouldFail: false},
        {actual: "hello", expected: "world",   shouldFail: true},
        {actual: "Hello", expected: "Helo",    shouldFail: true},
    }
    for _, test := range tests {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, test.actual, is.EqualToIgnoringCase(test.expected))
        then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
    }
}
```
