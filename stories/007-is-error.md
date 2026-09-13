# 007 — `is.Error` / `is.NoError` / `is.ErrorIs` / `is.ErrorAs`

**Status:** backlog  
**Package:** `is`  
**Hamcrest parity:** no (Go-specific; Java uses `assertThrows`)

## Summary

Go's explicit `error` return values are tested constantly but gocrest has no dedicated matchers for them.
Currently you write `is.Nil()` for no-error (loses the typed context) and roll your own string checks for error messages.

Four matchers:

- `is.NoError()` — passes when `error` is nil; clearer failure message than `is.Nil()`
- `is.AnError()` — passes when `error` is non-nil (any error)
- `is.ErrorWith(matcher)` — passes when `err.Error()` satisfies the string matcher
- `is.ErrorIs(target)` — passes when `errors.Is(actual, target)` is true (unwraps chain)
- `is.ErrorAs[T]()` — passes when `errors.As(actual, &T{})` succeeds (type in chain)

## API

```go
func NoError() *gocrest.Matcher[error]
func AnError() *gocrest.Matcher[error]
func ErrorWith(expected *gocrest.Matcher[string]) *gocrest.Matcher[error]
func ErrorIs(target error) *gocrest.Matcher[error]
func ErrorAs[T error]() *gocrest.Matcher[error]
```

## Example usage

```go
then.AssertThat(t, doThing(), is.NoError())
then.AssertThat(t, doThing(), is.ErrorWith(is.StringContaining("not found")))
then.AssertThat(t, doThing(), is.ErrorIs(os.ErrNotExist))
then.AssertThat(t, doThing(), is.ErrorAs[*os.PathError]())
```

## Acceptance criteria

### `is.NoError`
- Passes when `err == nil`
- Fails when `err != nil`
- Failure description: `no error`
- Actual in failure: `error: <err.Error()>`

### `is.AnError`
- Passes when `err != nil`
- Fails when `err == nil`
- Failure description: `an error`

### `is.ErrorWith`
- Passes when the error is non-nil and `err.Error()` satisfies the string matcher
- Fails when error is nil
- Fails when `err.Error()` does not satisfy the string matcher

### `is.ErrorIs`
- Passes when `errors.Is(actual, target)` returns true
- Fails otherwise
- Failure description: `error matching errors.Is(<target>)`

### `is.ErrorAs`
- Passes when `errors.As(actual, ...)` returns true for the type parameter
- Fails otherwise

## Test shape

```go
func TestNoError(t *testing.T) {
    then.AssertThat(t, error(nil), is.NoError())

    stubTestingT := new(StubTestingT)
    then.AssertThat(stubTestingT, fmt.Errorf("boom"), is.NoError())
    then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(true).Reason(stubTestingT.MockTestOutput))
}

func TestErrorWith(t *testing.T) {
    tests := []struct {
        actual     error
        matcher    *gocrest.Matcher[string]
        shouldFail bool
    }{
        {actual: fmt.Errorf("file not found"), matcher: is.StringContaining("not found"), shouldFail: false},
        {actual: fmt.Errorf("other error"),    matcher: is.StringContaining("not found"), shouldFail: true},
        {actual: nil,                          matcher: is.StringContaining("not found"), shouldFail: true},
    }
    for _, test := range tests {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, test.actual, is.ErrorWith(test.matcher))
        then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
    }
}
```
