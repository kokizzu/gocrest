# 005 — `is.CloseTo`

**Status:** backlog  
**Package:** `is`  
**Hamcrest parity:** yes (`closeTo(value, delta)`)

## Summary

Add `is.CloseTo(expected, delta)` for floating-point approximate equality.
`is.EqualTo` on floats is fragile due to rounding; this is a core Hamcrest matcher that gocrest is missing.

## API

```go
func CloseTo[A constraints.Float](expected, delta A) *gocrest.Matcher[A]
```

## Example usage

```go
then.AssertThat(t, 3.14159, is.CloseTo(3.14, 0.01))
then.AssertThat(t, float32(1.0/3.0), is.CloseTo(float32(0.333), float32(0.001)))
```

## Acceptance criteria

- Passes when `math.Abs(actual - expected) <= delta`
- Fails when the difference exceeds delta
- Works for `float32` and `float64`
- Failure description: `a value within <delta> of <expected>`
- Actual in failure: `<actual> differed by <diff>`
- Panics (or fails clearly) if delta is negative

## Test shape

```go
func TestCloseTo(t *testing.T) {
    tests := []struct {
        actual     float64
        expected   float64
        delta      float64
        shouldFail bool
    }{
        {actual: 3.14159, expected: 3.14,  delta: 0.01,  shouldFail: false},
        {actual: 3.14159, expected: 3.14,  delta: 0.001, shouldFail: true},
        {actual: 1.0,     expected: 1.0,   delta: 0.0,   shouldFail: false},
        {actual: 1.1,     expected: 1.0,   delta: 0.05,  shouldFail: true},
        {actual: -1.0,    expected: -1.01, delta: 0.02,  shouldFail: false},
    }
    for _, test := range tests {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, test.actual, is.CloseTo(test.expected, test.delta))
        then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
    }
}
```
