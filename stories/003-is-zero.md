# 003 — `is.Zero`

**Status:** backlog  
**Package:** `is`  
**Hamcrest parity:** no (Go-specific; Hamcrest uses `equalTo(0)` or `nullValue`)

## Summary

Add `is.Zero[T]()` that passes when actual equals the zero value for its type (`0`, `""`, `false`, `nil` pointer, zero struct, etc.).
Avoids the awkward `is.EqualTo("")` or `is.EqualTo(0)` for cases where the concrete zero is unknown at the call site (e.g. a generic helper).

## API

```go
func Zero[A comparable]() *gocrest.Matcher[A]
```

## Example usage

```go
then.AssertThat(t, 0, is.Zero[int]())
then.AssertThat(t, "", is.Zero[string]())
then.AssertThat(t, false, is.Zero[bool]())
```

## Acceptance criteria

- Passes when actual is the zero value of its type
- Fails when actual is non-zero
- Works for any `comparable` type
- Failure description: `zero value`
- Actual in failure: `<actual>`

## Test shape

```go
func TestIsZero(t *testing.T) {
    t.Run("int zero passes", func(t *testing.T) {
        then.AssertThat(t, 0, is.Zero[int]())
    })
    t.Run("int non-zero fails", func(t *testing.T) {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, 1, is.Zero[int]())
        then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(true).Reason(stubTestingT.MockTestOutput))
    })
    t.Run("string zero passes", func(t *testing.T) {
        then.AssertThat(t, "", is.Zero[string]())
    })
    t.Run("bool zero passes", func(t *testing.T) {
        then.AssertThat(t, false, is.Zero[bool]())
    })
}
```
