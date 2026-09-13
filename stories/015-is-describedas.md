# 015 — `is.DescribedAs`

**Status:** backlog  
**Package:** `is`  
**Hamcrest parity:** yes (`describedAs`)

## Summary

Add `is.DescribedAs(description, matcher)` — a decorator that wraps any matcher and replaces its failure description with a custom string.

Useful when the generated description is accurate but not readable in context, without needing to write a whole custom matcher.

## API

```go
func DescribedAs[A any](description string, matcher *gocrest.Matcher[A]) *gocrest.Matcher[A]
```

## Example usage

```go
then.AssertThat(t, user.Role,
    is.DescribedAs("user must be an admin", is.EqualTo("admin")))

then.AssertThat(t, response.StatusCode,
    is.DescribedAs("expected HTTP 200 OK", is.EqualTo(200)))
```

## Acceptance criteria

- Delegates `Matches` to the wrapped matcher unchanged
- Replaces `Describe` with the supplied description string
- Preserves `Actual` from the wrapped matcher on failure
- The wrapped matcher's own description does not appear in the failure output
- Failure output contains exactly the supplied description

## Test shape

```go
func TestDescribedAs(t *testing.T) {
    t.Run("passes when inner matcher passes", func(t *testing.T) {
        then.AssertThat(t, 1, is.DescribedAs("should be one", is.EqualTo(1)))
    })

    t.Run("fails when inner matcher fails", func(t *testing.T) {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, 2, is.DescribedAs("should be one", is.EqualTo(1)))
        then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(true))
    })

    t.Run("failure output uses custom description not inner description", func(t *testing.T) {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, 2, is.DescribedAs("should be one", is.EqualTo(1)))
        then.AssertThat(t, stubTestingT.MockTestOutput, is.AllOf(
            is.StringContaining("should be one"),
            is.Not(is.StringContaining("value equal to")),
        ))
    })
}
```
