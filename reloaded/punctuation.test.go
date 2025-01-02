package reloaded

import (
	"testing"
)

func TestPunctuationRegEx(t *testing.T) {

}
func TestPunctuationRegExHandlesMultipleSpacesAndTabs(t *testing.T) {
	input := "    \t  Hello, world!   \t  "
	expected := "Hello, world!"
	result := PunctuationRegEx(input)
	if result != expected {
		t.Errorf("PunctuationRegEx(%q) = %q, want %q", input, result, expected)
	}
}

func TestPunctuationRegExMultipleSequentialPunctuation(t *testing.T) {
	input := "Hello ... world ! ? How are you ?"
	expected := "Hello... world!? How are you?"
	result := PunctuationRegEx(input)
	if result != expected {
		t.Errorf("PunctuationRegEx(%q) = %q, want %q", input, result, expected)
	}
}


func TestPunctuationRegExRemovesSpacesBeforePunctuation(t *testing.T) {
	input := "Hello   , world   ! How are you   ?"
	expected := "Hello, world! How are you?"
	result := PunctuationRegEx(input)
	if result != expected {
		t.Errorf("PunctuationRegEx(%q) = %q, want %q", input, result, expected)
	}
}

func TestPunctuationRegExHandlesColonsBeforeQuotes(t *testing.T) {
	input := "He said (quietly): 'Hello' and then (loudly): 'Goodbye!'"
	expected := "He said (quietly): 'Hello' and then (loudly): 'Goodbye!'"
	result := PunctuationRegEx(input)
	if result != expected {
		t.Errorf("PunctuationRegEx(%q) = %q, want %q", input, result, expected)
	}
}

func TestPunctuationRegExCleanupSpacesInsideQuotes(t *testing.T) {
	input := "He said: ' This   is  a   test   with    spaces   ' and continued."
	expected := "He said: 'This is a test with spaces' and continued."
	result := PunctuationRegEx(input)
	if result != expected {
		t.Errorf("PunctuationRegEx(%q) = %q, want %q", input, result, expected)
	}
}

func TestPunctuationRegExFixesQuotedText(t *testing.T) {
	input := "He said: ' Hello , world ! How are you ? '"
	expected := "He said: 'Hello, world! How are you?'"
	result := PunctuationRegEx(input)
	if result != expected {
		t.Errorf("PunctuationRegEx(%q) = %q, want %q", input, result, expected)
	}
}

func TestPunctuationRegExHandlesOnlyPunctuation(t *testing.T) {
	input := "... !? , . : ;"
	expected := "...!?,.:"
	result := PunctuationRegEx(input)
	if result != expected {
		t.Errorf("PunctuationRegEx(%q) = %q, want %q", input, result, expected)
	}
}

func TestPunctuationRegExCleanUpMultipleSpaces(t *testing.T) {
	input := "This   is  a   test   with    multiple    spaces."
	expected := "This is a test with multiple spaces."
	result := PunctuationRegEx(input)
	if result != expected {
		t.Errorf("PunctuationRegEx(%q) = %q, want %q", input, result, expected)
	}
}

func TestPunctuationRegExMixedQuotedUnquotedText(t *testing.T) {
	input := "He said: ' Hello , world ! 'Then he paused . ' How are you ? ' he asked ."
	expected := "He said: 'Hello, world!' Then he paused. 'How are you?' he asked."
	result := PunctuationRegEx(input)
	if result != expected {
		t.Errorf("PunctuationRegEx(%q) = %q, want %q", input, result, expected)
	}
}
