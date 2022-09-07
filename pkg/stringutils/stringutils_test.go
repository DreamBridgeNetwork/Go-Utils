package stringutils

import "testing"

func TestReverse(t *testing.T) {
	want := "leafaR"

	got := Reverse("Rafael")

	if got != want {
		t.Errorf("Reverse() = %q, want %q", got, want)
	}
}
