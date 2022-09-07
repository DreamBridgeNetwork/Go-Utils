package password

import "testing"

func TestNewPassword(t *testing.T) {
	want := 16
	got := NewPassword(want)

	if len(got) != want {
		t.Error("Reverse() = Wanted size: ", want, ", size got: ", len(got))
	}

	t.Log("NewPassword() = ", got)
}
