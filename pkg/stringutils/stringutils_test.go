package stringutils

import (
	"log"
	"testing"
)

func TestReverse(t *testing.T) {
	want := "leafaR"

	got := Reverse("Rafael")

	if got != want {
		t.Errorf("Reverse() = %q, want %q", got, want)
	}
}

func TestGenerateStringHash(t *testing.T) {
	log.Println("TestGenerateStringHash")

	password := "This is my password"

	hash, err := GenerateStringHash(password)
	if err != nil {
		t.Error("Error generating hash: ", err)
		return
	}

	log.Println("Hash: ", hash)

	resp, err := CheckStringHash(password, hash)

	if err != nil {
		t.Error("Error testing if hash matchs string: ", err)
		return
	}

	if !resp {
		t.Error("Hash didn´t match string.")
		return
	}

	log.Println("TestGenerateStringHash OK")
}
