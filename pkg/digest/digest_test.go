package digest

import "testing"

func TestGenerateDigestSHA256(t *testing.T) {
	want := "XO+A3vP1LxWUJyDzuf+xrIFa4I0m+F+0L3/gkc0XrXA="

	got, err := GenerateDigestSHA256("Teste Digest.")

	if err != nil {
		t.Error("GenerateDigestSHA256() = ", err.Error())
	}

	if got != want {
		t.Errorf("GenerateDigestSHA256() = %q, want %q", got, want)
	}
}
