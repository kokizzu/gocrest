# 011 — `is.ArrayContainingExactly`

**Status:** backlog  
**Package:** `is`  
**Hamcrest parity:** yes (`containsInAnyOrder`)

## Summary

Add `is.ArrayContainingExactly(vals...)` — passes when the actual slice contains exactly the given elements in any order, with no extras.

This is the unordered exact-set counterpart to story 010.

The existing `is.ArrayContaining` allows extras; this does not.

## API

```go
func ArrayContainingExactly[A comparable](expected ...A) *gocrest.Matcher[[]A]
```

## Example usage

```go
then.AssertThat(t, []string{"c", "a", "b"}, is.ArrayContainingExactly("a", "b", "c"))
```

## Acceptance criteria

- Passes when actual contains exactly the expected elements (any order, duplicates must match)
- Fails when actual has extra elements
- Fails when actual is missing elements
- Fails when lengths differ
- Failure description: `a slice containing exactly [<a> <b> <c>] in any order`

## Test shape

```go
func TestArrayContainingExactly(t *testing.T) {
    tests := []struct {
        actual     []string
        expected   []string
        shouldFail bool
    }{
        {actual: []string{"c", "a", "b"}, expected: []string{"a", "b", "c"}, shouldFail: false},
        {actual: []string{"a", "b", "c"}, expected: []string{"a", "b", "c"}, shouldFail: false},
        {actual: []string{"a", "b"},      expected: []string{"a", "b", "c"}, shouldFail: true},
        {actual: []string{"a", "b", "c", "d"}, expected: []string{"a", "b", "c"}, shouldFail: true},
        {actual: []string{"a", "a", "b"}, expected: []string{"a", "b", "b"}, shouldFail: true},
    }
    for _, test := range tests {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, test.actual, is.ArrayContainingExactly(test.expected...))
        then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
    }
}
```
