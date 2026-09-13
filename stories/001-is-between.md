# 001 — `is.Between`

**Status:** backlog  
**Package:** `is`  
**Hamcrest parity:** yes (`between` / `allOf(greaterThanOrEqualTo, lessThanOrEqualTo)`)

## Summary

Add `is.Between(lo, hi)` that passes when `lo <= actual <= hi`.
Currently this requires `is.AllOf(is.GreaterThanOrEqualTo(lo), is.LessThanOrEqualTo(hi))` which is verbose and produces a poor failure message.

## API

```go
func Between[A constraints.Ordered](lo, hi A) *gocrest.Matcher[A]
```

## Example usage

```go
then.AssertThat(t, 5, is.Between(1, 10))
then.AssertThat(t, "mango", is.Between("apple", "orange"))
```

## Acceptance criteria

- Passes when `actual == lo`
- Passes when `actual == hi`
- Passes when `lo < actual < hi`
- Fails when `actual < lo`
- Fails when `actual > hi`
- Works for all `constraints.Ordered` types (int, float64, string, …)
- Failure description reads: `value between <lo> and <hi>`
- Actual in failure reads: `<actual>`

## Test shape

```go
func TestBetween(t *testing.T) {
    tests := []struct {
        actual     int
        lo         int
        hi         int
        shouldFail bool
    }{
        {actual: 5,  lo: 1,  hi: 10, shouldFail: false},
        {actual: 1,  lo: 1,  hi: 10, shouldFail: false},
        {actual: 10, lo: 1,  hi: 10, shouldFail: false},
        {actual: 0,  lo: 1,  hi: 10, shouldFail: true},
        {actual: 11, lo: 1,  hi: 10, shouldFail: true},
    }
    for _, test := range tests {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, test.actual, is.Between(test.lo, test.hi))
        then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
    }
}
```
