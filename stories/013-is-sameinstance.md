# 013 — `is.SameInstance`

**Status:** backlog  
**Package:** `is`  
**Hamcrest parity:** yes (`sameInstance`)

## Summary

Add `is.SameInstance(expected)` — passes when actual and expected are the same pointer (reference equality).

`is.EqualTo` uses `reflect.DeepEqual` and passes for two distinct pointers with the same contents.
This matcher uses `actual == expected` (pointer identity) to verify same object, not same value.

## API

```go
func SameInstance[A comparable](expected A) *gocrest.Matcher[A]
```

## Example usage

```go
x := &MyStruct{Value: 1}
y := x
z := &MyStruct{Value: 1}

then.AssertThat(t, y, is.SameInstance(x))           // passes — same pointer
then.AssertThat(stubT, z, is.SameInstance(x))       // fails — different pointer, same value
```

## Acceptance criteria

- Passes when `actual == expected` (pointer identity)
- Fails when `actual != expected` even if `reflect.DeepEqual` would return true
- Works for any `comparable` type
- Failure description: `the same instance as <expected>`
- Actual in failure: `<actual>`

## Test shape

```go
func TestSameInstance(t *testing.T) {
    type T struct{ V int }
    x := &T{V: 1}
    y := x
    z := &T{V: 1}

    then.AssertThat(t, y, is.SameInstance(x))

    stubTestingT := new(StubTestingT)
    then.AssertThat(stubTestingT, z, is.SameInstance(x))
    then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(true).Reason(stubTestingT.MockTestOutput))
}
```
