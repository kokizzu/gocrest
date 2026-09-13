# 012 — `has.Value` / `has.Entry`

**Status:** backlog  
**Package:** `has`  
**Hamcrest parity:** yes (`hasValue`, `hasEntry`)

## Summary

Two map matchers that complete the set alongside `has.Key` / `has.AllKeys`:

- `has.Value(v)` — map contains this value at any key
- `has.Entry(k, v)` — map contains this exact key→value pair

## API

```go
func Value[K comparable, V comparable](expected V) *gocrest.Matcher[map[K]V]
func Entry[K comparable, V comparable](key K, value V) *gocrest.Matcher[map[K]V]
```

## Example usage

```go
then.AssertThat(t, map[string]int{"a": 1, "b": 2}, has.Value[string, int](2))
then.AssertThat(t, map[string]int{"a": 1, "b": 2}, has.Entry("a", 1))
```

## Acceptance criteria

### `has.Value`
- Passes when any map value equals expected
- Fails when no value matches
- Failure description: `map has value '<expected>'`

### `has.Entry`
- Passes when `actual[key] == value`
- Fails when key is absent
- Fails when key exists but value differs
- Failure description: `map has entry {'<key>': '<value>'}`

## Test shape

```go
func TestHasValue(t *testing.T) {
    tests := []struct {
        actual     map[string]int
        expected   int
        shouldFail bool
    }{
        {actual: map[string]int{"a": 1, "b": 2}, expected: 2,  shouldFail: false},
        {actual: map[string]int{"a": 1, "b": 2}, expected: 99, shouldFail: true},
        {actual: map[string]int{},                expected: 1,  shouldFail: true},
    }
    for _, test := range tests {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, test.actual, has.Value[string, int](test.expected))
        then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
    }
}

func TestHasEntry(t *testing.T) {
    tests := []struct {
        actual     map[string]int
        key        string
        value      int
        shouldFail bool
    }{
        {actual: map[string]int{"a": 1}, key: "a", value: 1,  shouldFail: false},
        {actual: map[string]int{"a": 1}, key: "a", value: 2,  shouldFail: true},
        {actual: map[string]int{"a": 1}, key: "b", value: 1,  shouldFail: true},
    }
    for _, test := range tests {
        stubTestingT := new(StubTestingT)
        then.AssertThat(stubTestingT, test.actual, has.Entry(test.key, test.value))
        then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
    }
}
```
