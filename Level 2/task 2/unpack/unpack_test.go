package unpack

import "testing"

func TestIsEscapeSymbol(t *testing.T) {
	validEscapeRunes := []rune(`\`)
	for _, escapeRune := range validEscapeRunes {
		if !isEscape(escapeRune) {
			t.Errorf("Symbol \"%s\" [rune %d] is not escape symbol", string(escapeRune), escapeRune)
		}
	}

	invalidEscapeRunes := []rune(`|/#@!`)
	for _, escapeRune := range invalidEscapeRunes {
		if isEscape(escapeRune) {
			t.Errorf("Symbol \"%s\" [rune %d] is valid escape symbol", string(escapeRune), escapeRune)
		}
	}
}

func TestCaseRegular(t *testing.T) {
	data := "a4bc2d5e"
	expect := "aaaabccddddde"

	if result, _ := UnpackingString(data); result != expect {
		t.Errorf("Invalid result: expect \"%s\", got \"%s\"", expect, result)
	}
}

func TestCaseSimple(t *testing.T) {
	data := "abcd"
	expect := "abcd"

	if result, _ := UnpackingString(data); result != expect {
		t.Errorf("Invalid result: expect \"%s\", got \"%s\"", expect, result)
	}
}

func TestCaseIncorrectOnlyNumbers(t *testing.T) {
	data := "45"

	if result, err := UnpackingString(data); err == nil {
		t.Errorf("Passed invalid data with result \"%s\", expect error", result)
	}
}

func TestCaseEmptyString(t *testing.T) {
	if result, _ := UnpackingString(""); result != "" {
		t.Errorf("Invalid result: expect empty string, got \"%s\"", result)
	}
}

func TestCaseIncorrectMissedLetter(t *testing.T) {
	data := "a2b37"

	if result, err := UnpackingString(data); err == nil {
		t.Errorf("Passed invalid data with result \"%s\", expect error", result)
	}
}

func TestCaseEscapedNumbers(t *testing.T) {
	data := `qwe\4\5`
	expect := "qwe45"

	if result, _ := UnpackingString(data); result != expect {
		t.Errorf("Invalid result: expect \"%s\", got \"%s\"", expect, result)
	}
}

func TestCasePackedNumber(t *testing.T) {
	data := `qwe\45`
	expect := "qwe44444"

	if result, _ := UnpackingString(data); result != expect {
		t.Errorf("Invalid result: expect \"%s\", got \"%s\"", expect, result)
	}
}

func TestCasePackedSlash(t *testing.T) {
	data := `qwe\\5`
	expect := `qwe\\\\\`

	if result, _ := UnpackingString(data); result != expect {
		t.Errorf("Invalid result: expect \"%s\", got \"%s\"", expect, result)
	}
}

func TestCaseComplexSymbols(t *testing.T) {
	data := "▬2♠3♫"
	expect := "▬▬♠♠♠♫"

	if result, _ := UnpackingString(data); result != expect {
		t.Errorf("Invalid result: expect \"%s\", got \"%s\"", expect, result)
	}
}
