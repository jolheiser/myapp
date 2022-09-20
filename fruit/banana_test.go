package fruit

import "testing"

// A function inside a *_test.go file that starts with Test* and takes a (*testing.T) is the most basic way to add a Go test
func TestBanana(t *testing.T) {
	const expected = "banana"

	b := Banana{}

	got := b.String()

	if got != expected {
		t.Logf("expected `banana` but got %s", got)
		t.Fail()
	}
}
