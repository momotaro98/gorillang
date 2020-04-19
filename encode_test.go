package gorillang

import (
	"testing"
)

func TestEncode(t *testing.T) {
	// Arrange
	input := `おはよう日本`
	expected := `ウホゥホウホホゥゥゥ ウホゥホウホウォウッホ ウホゥホウホ？？ ウホゥホウホホゥウォ ウホウォホッ！！ホッ ウホウォうほホホッ！`
	// Act
	act := Encode(input)
	// Assert
	if act != expected {
		t.Errorf("wrong. expected=%s, got=%s", expected, act)
	}
}
