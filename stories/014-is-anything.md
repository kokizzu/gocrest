# 014 — `is.Anything`

**Status:** backlog  
**Package:** `is`  
**Hamcrest parity:** yes (`anything()`)

## Summary

Add `is.Anything[T]()` — a matcher that always passes regardless of the actual value.

Primary use: placeholder in composite matchers where one position is "don't care".

```go
then.AssertThat(t, myMap, has.Entry(is.Anything[string](), is.EqualTo("foo")))
// (hypothetical — needs matcher-accepting Entry variant)
```

Also useful in `has.EveryElement` when some positions should be unchecked.

## API

```go
func Anything[A any]() *gocrest.Matcher[A]
```

## Example usage

```go
then.AssertThat(t, "whatever", is.Anything[string]())
then.AssertThat(t, 42, is.Anything[int]())
then.AssertThat(t, (*MyStruct)(nil), is.Anything[*MyStruct]())
```

## Acceptance criteria

- Always passes regardless of actual value (including nil, zero, any pointer)
- Failure description: `anything`
- Never sets Actual (it never fails)

## Test shape

```go
func TestAnything(t *testing.T) {
    then.AssertThat(t, "any string", is.Anything[string]())
    then.AssertThat(t, 0, is.Anything[int]())
    then.AssertThat(t, (*testing.T)(nil), is.Anything[*testing.T]())

    then.AssertThat(t, []string{"a", "b", "c"}, has.EveryElement(
        is.Anything[string](),
        is.Anything[string](),
        is.Anything[string](),
    ))
}
```
