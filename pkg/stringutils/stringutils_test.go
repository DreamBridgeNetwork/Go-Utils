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

func TestValidateStringRegex(t *testing.T) {
	log.Println("TestValidateStringRegex")

	// Positive test
	regexExpr := "^[1-9]\\d*$" // Integer number

	data := "1234"

	resp, err := ValidateStringRegex(data, regexExpr)
	if err != nil {
		t.Error("Error validating regex in the positive test: ", err)
	}

	if !resp {
		t.Error("Error validating data in the positive test.")
	}

	// Negative tests
	data = "1234A"

	resp, err = ValidateStringRegex(data, regexExpr)
	if err != nil {
		t.Error("Error validating regex in the negative test: ", err)
	}

	if resp {
		t.Error("Error invalidating data in the negative test.")
	}

	// Error test

	regexExpr = "?[1-9]\\d*$" // Integer number
	_, err = ValidateStringRegex(data, regexExpr)
	if err == nil {
		t.Error("Error in the error test. No error returned.")
	}

	log.Println("TestValidateStringRegex OK")
}
