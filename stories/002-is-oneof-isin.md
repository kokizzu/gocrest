# 002 — `is.OneOf` / `is.IsIn`

**Status:** backlog  
**Package:** `is`  
**Hamcrest parity:** yes (`oneOf`, `isIn`)

## Summary

Two membership matchers:

- `is.OneOf(vals...)` — actual equals one of the supplied values (equivalent to `is.AnyOf(is.EqualTo(a), is.EqualTo(b), ...)` but terser)
- `is.IsIn(collection)` — actual is a member of the supplied slice (inverse perspective of `is.ArrayContaining`)

## API

```go
func OneOf[A comparable](vals ...A) *gocrest.Matcher[A]
func IsIn[A comparable](collection []A) *gocrest.Matcher[A]
```

## Example usage

```go
then.AssertThat(t, "red", is.OneOf("red", "green", "blue"))
then.AssertThat(t, 3, is.IsIn([]int{1, 2, 3, 4, 5}))
```

## Acceptance criteria

### `is.OneOf`
- Passes when actual equals any of the supplied values
- Fails when actual matches none
- Failure description: `one of [<a> <b> ...]`

### `is.IsIn`
- Passes when actual is present in the collection
- Fails when actual is absent
- Failure description: `a value in [<a> <b> ...]`

## Test shape

```go
func TestOneOf(t *testing.T) {
    tests := []struct {
        actual     string
        vals       []string
        shouldFail bool
    }{
        {actual: "red",    vals: []string{"red", "green", "blue"}, shouldFail: false},
        {actual: "yellow", vals: []string{"red", "green", "blue"}, shouldFail: true},
    }
    for _, test := range tests {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, test.actual, is.OneOf(test.vals...))
        then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
    }
}

func TestIsIn(t *testing.T) {
    tests := []struct {
        actual     int
        collection []int
        shouldFail bool
    }{
        {actual: 3, collection: []int{1, 2, 3}, shouldFail: false},
        {actual: 9, collection: []int{1, 2, 3}, shouldFail: true},
    }
    for _, test := range tests {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, test.actual, is.IsIn(test.collection))
        then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
    }
}
```
