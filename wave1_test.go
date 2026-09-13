package gocrest_test

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
)

// --- is.Anything ---

func TestAnythingAlwaysPasses(t *testing.T) {
	then.AssertThat(t, "any string", is.Anything[string]())
	then.AssertThat(t, 0, is.Anything[int]())
	then.AssertThat(t, (*testing.T)(nil), is.Anything[*testing.T]())
}

func TestAnythingDescription(t *testing.T) {
	then.AssertThat(t, is.Anything[string]().Describe, is.EqualTo("anything"))
}

// --- is.SameInstance ---

func TestSameInstancePasses(t *testing.T) {
	type T struct{ V int }
	x := &T{V: 1}
	y := x
	then.AssertThat(t, y, is.SameInstance(x))
}

func TestSameInstanceFailsForDifferentPointer(t *testing.T) {
	type T struct{ V int }
	x := &T{V: 1}
	z := &T{V: 1}
	stubTestingT := new(StubTestingT)
	then.AssertThat(stubTestingT, z, is.SameInstance(x))
	then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(true).Reason(stubTestingT.MockTestOutput))
}

func TestSameInstanceDescription(t *testing.T) {
	type T struct{ V int }
	x := &T{V: 1}
	stubTestingT := new(StubTestingT)
	then.AssertThat(stubTestingT, &T{V: 2}, is.SameInstance(x))
	then.AssertThat(t, stubTestingT.MockTestOutput, is.StringContaining("the same instance as"))
}

// --- is.Zero ---

func TestZero(t *testing.T) {
	tests := []struct {
		name       string
		run        func(*StubTestingT)
		shouldFail bool
	}{
		{
			name:       "int zero passes",
			run:        func(st *StubTestingT) { then.AssertThat(st, 0, is.Zero[int]()) },
			shouldFail: false,
		},
		{
			name:       "int non-zero fails",
			run:        func(st *StubTestingT) { then.AssertThat(st, 1, is.Zero[int]()) },
			shouldFail: true,
		},
		{
			name:       "string zero passes",
			run:        func(st *StubTestingT) { then.AssertThat(st, "", is.Zero[string]()) },
			shouldFail: false,
		},
		{
			name:       "string non-zero fails",
			run:        func(st *StubTestingT) { then.AssertThat(st, "x", is.Zero[string]()) },
			shouldFail: true,
		},
		{
			name:       "bool zero passes",
			run:        func(st *StubTestingT) { then.AssertThat(st, false, is.Zero[bool]()) },
			shouldFail: false,
		},
		{
			name:       "bool non-zero fails",
			run:        func(st *StubTestingT) { then.AssertThat(st, true, is.Zero[bool]()) },
			shouldFail: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stubTestingT := new(StubTestingT)
			test.run(stubTestingT)
			then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
		})
	}
}

func TestZeroDescription(t *testing.T) {
	stubTestingT := new(StubTestingT)
	then.AssertThat(stubTestingT, 1, is.Zero[int]())
	then.AssertThat(t, stubTestingT.MockTestOutput, is.StringContaining("zero value"))
}

type zeroTestErr struct{ msg string }

func (e *zeroTestErr) Error() string { return e.msg }

func TestZeroTypedNilInterfaceFails(t *testing.T) {
	// A typed-nil ((*zeroTestErr)(nil) stored as error) is not a zero interface —
	// the interface itself is non-nil. Zero[error]() must fail for this value.
	var typedNil error = (*zeroTestErr)(nil)
	stubTestingT := new(StubTestingT)
	then.AssertThat(stubTestingT, typedNil, is.Zero[error]())
	then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(true).Reason("typed-nil interface should not be zero"))
}

func TestSameInstanceActualClearedOnPass(t *testing.T) {
	// Verifies stale Actual is not left over after a pass — important when used inside AllOf.
	type T struct{ V int }
	x := &T{V: 1}
	m := is.SameInstance(x)
	z := &T{V: 2}
	m.Matches(z) // fail — sets Actual
	m.Matches(x) // pass — must clear Actual
	then.AssertThat(t, m.Actual, is.EqualTo(""))
}

func TestZeroActualClearedOnPass(t *testing.T) {
	m := is.Zero[int]()
	m.Matches(1) // fail — sets Actual
	m.Matches(0) // pass — must clear Actual
	then.AssertThat(t, m.Actual, is.EqualTo(""))
}
