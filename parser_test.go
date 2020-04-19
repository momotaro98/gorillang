package gorillang

import (
	"fmt"
	"testing"
)

func TestSentence(t *testing.T) {
	// Arrange
	// こんにちは
	input := `ウホゥホウホホッゥホ ウホゥホウホオホゥホ ウホゥホウホウォッッ ウホゥホウホウォウウ ウホゥホウホウォウッホ`

	tests := []struct {
		expectedCodePoint CodePoint
	}{
		{"3053"}, // こ
		{"3093"}, // ん
		{"306B"}, // に
		{"3061"}, // ち
		{"306F"}, // は
	}

	l := NewLexer(input)
	// Act
	p := NewParser(l)

	// Assert
	sentence := p.ParseSentence()
	if sentence == nil {
		t.Fatalf("Sentence() returned nil")
	}
	for i, tt := range tests {
		if sentence.CodePoints[i] != tt.expectedCodePoint {
			t.Errorf("[%d] wrong. expected=%+v, got=%+v", i, tt.expectedCodePoint, sentence.CodePoints[i])
		}
	}
	// Act and Assert String() method
	act := fmt.Sprintf("%s", sentence)
	if act != "こんにちは" {
		t.Errorf("String() has wrong. expected=こんにちは, got=%s", act)
	}
}
