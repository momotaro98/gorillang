package gorillang

import (
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	// Arrange
	var (
		inAtABC = `ウホホゥウホ ウホホゥウウ ウホホゥホホ ウホホゥゥホ`  // @ABC
		inDEFG  = `ウホホゥホゥ ウホホゥホッ ウホホゥウォ ウホホゥうほ`  // DEFG
		inHIJK  = `ウホホゥ？ ウホホゥオホ ウホホゥゥゥ ウホホゥッッ`   // HIJK
		inLMNO  = `ウホホゥッ！ ウホホゥォ！ ウホホゥ！！ ウホホゥウッホ` // LMNO
	)
	input := strings.Join([]string{inAtABC, inDEFG, inHIJK, inLMNO}, ` `)

	tests := []struct {
		expectedType    TokenType
		expectedLiteral []rune
	}{
		// @ABC
		{expectedType: PREFIX, expectedLiteral: []rune(`ウホ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: X0, expectedLiteral: []rune(`ウホ`)},
		{expectedType: WHITESPACE, expectedLiteral: []rune(` `)},
		{expectedType: PREFIX, expectedLiteral: []rune(`ウホ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: X1, expectedLiteral: []rune(`ウウ`)},
		{expectedType: WHITESPACE, expectedLiteral: []rune(` `)},
		{expectedType: PREFIX, expectedLiteral: []rune(`ウホ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: X2, expectedLiteral: []rune(`ホホ`)},
		{expectedType: WHITESPACE, expectedLiteral: []rune(` `)},
		{expectedType: PREFIX, expectedLiteral: []rune(`ウホ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: X3, expectedLiteral: []rune(`ゥホ`)},
		{expectedType: WHITESPACE, expectedLiteral: []rune(` `)},
		// DEFG
		{expectedType: PREFIX, expectedLiteral: []rune(`ウホ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: WHITESPACE, expectedLiteral: []rune(` `)},
		{expectedType: PREFIX, expectedLiteral: []rune(`ウホ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: X5, expectedLiteral: []rune(`ホッ`)},
		{expectedType: WHITESPACE, expectedLiteral: []rune(` `)},
		{expectedType: PREFIX, expectedLiteral: []rune(`ウホ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: X6, expectedLiteral: []rune(`ウォ`)},
		{expectedType: WHITESPACE, expectedLiteral: []rune(` `)},
		{expectedType: PREFIX, expectedLiteral: []rune(`ウホ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: X7, expectedLiteral: []rune(`うほ`)},
		{expectedType: WHITESPACE, expectedLiteral: []rune(` `)},
		// HIJK
		{expectedType: PREFIX, expectedLiteral: []rune(`ウホ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: X8, expectedLiteral: []rune(`？`)},
		{expectedType: WHITESPACE, expectedLiteral: []rune(` `)},
		{expectedType: PREFIX, expectedLiteral: []rune(`ウホ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: X9, expectedLiteral: []rune(`オホ`)},
		{expectedType: WHITESPACE, expectedLiteral: []rune(` `)},
		{expectedType: PREFIX, expectedLiteral: []rune(`ウホ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: XA, expectedLiteral: []rune(`ゥゥ`)},
		{expectedType: WHITESPACE, expectedLiteral: []rune(` `)},
		{expectedType: PREFIX, expectedLiteral: []rune(`ウホ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: XB, expectedLiteral: []rune(`ッッ`)},
		{expectedType: WHITESPACE, expectedLiteral: []rune(` `)},
		// LMNO
		{expectedType: PREFIX, expectedLiteral: []rune(`ウホ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: XC, expectedLiteral: []rune(`ッ！`)},
		{expectedType: WHITESPACE, expectedLiteral: []rune(` `)},
		{expectedType: PREFIX, expectedLiteral: []rune(`ウホ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: XD, expectedLiteral: []rune(`ォ！`)},
		{expectedType: WHITESPACE, expectedLiteral: []rune(` `)},
		{expectedType: PREFIX, expectedLiteral: []rune(`ウホ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: XE, expectedLiteral: []rune(`！！`)},
		{expectedType: WHITESPACE, expectedLiteral: []rune(` `)},
		{expectedType: PREFIX, expectedLiteral: []rune(`ウホ`)},
		{expectedType: X4, expectedLiteral: []rune(`ホゥ`)},
		{expectedType: XF, expectedLiteral: []rune(`ウッホ`)},
		{expectedType: EOF, expectedLiteral: []rune{0}},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		// Act
		tok := l.NextToken()

		// Assert
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%+v, got=%+v", i, tt.expectedType, tok.Type)
		}

		for i := 0; i < len(tok.Literal); i++ {
			if tok.Literal[i] != tt.expectedLiteral[i] {
				t.Fatalf("tests[%d] - literal wrong. expected=%+v, got=%+v", i, tt.expectedLiteral, tok.Literal)
			}
		}
	}
}
