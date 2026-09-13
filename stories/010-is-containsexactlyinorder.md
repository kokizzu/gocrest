# 010 — `is.ArrayContainingExactlyInOrder`

**Status:** backlog  
**Package:** `is`  
**Hamcrest parity:** yes (`contains(a, b, c)`)

## Summary

Add `is.ArrayContainingExactlyInOrder(vals...)` — passes when the actual slice contains exactly the given elements in exactly the given order (no more, no less).

Existing matchers cover:
- `is.ArrayContaining` — these elements are present, any order, extras allowed
- `has.EveryElement` — each position matches a matcher (positional, exact length)

This adds: exact values, exact order, exact length — the simplest "is this slice exactly X?" check.

## API

```go
func ArrayContainingExactlyInOrder[A comparable](expected ...A) *gocrest.Matcher[[]A]
```

## Example usage

```go
then.AssertThat(t, []string{"a", "b", "c"}, is.ArrayContainingExactlyInOrder("a", "b", "c"))
```

## Acceptance criteria

- Passes when actual matches expected element-for-element
- Fails when lengths differ
- Fails when any element differs
- Fails when elements are in different order
- Failure description: `a slice containing exactly [<a> <b> <c>] in order`

## Test shape

```go
func TestArrayContainingExactlyInOrder(t *testing.T) {
    tests := []struct {
        actual     []string
        expected   []string
        shouldFail bool
    }{
        {actual: []string{"a", "b", "c"}, expected: []string{"a", "b", "c"}, shouldFail: false},
        {actual: []string{"a", "c", "b"}, expected: []string{"a", "b", "c"}, shouldFail: true},
        {actual: []string{"a", "b"},      expected: []string{"a", "b", "c"}, shouldFail: true},
        {actual: []string{"a", "b", "c", "d"}, expected: []string{"a", "b", "c"}, shouldFail: true},
        {actual: []string{},              expected: []string{},              shouldFail: false},
    }
    for _, test := range tests {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, test.actual, is.ArrayContainingExactlyInOrder(test.expected...))
        then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
    }
}
```
