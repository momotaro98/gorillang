package gorillang

type Lexer struct {
	input             []rune
	position          int  // 入力における現在の位置(現在の文字を指し示す)
	readPosition      int  // これから読み込む位置(現在の文字の次)
	ch                rune // 現在検査中の文字
	isAmongWhiteSpace bool // ホワイトスペースの間で文字(Unicode)の途中であるかどうか
}

func NewLexer(input string) *Lexer {
	r := []rune(input)
	l := &Lexer{
		input: r,
	}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	var tok Token

	take2 := func(tt TokenType) Token {
		ch := l.ch
		l.readChar()
		literal := []rune{ch, l.ch}
		return Token{Type: tt, Literal: literal}
	}

	switch l.ch {
	case []rune(` `)[0]:
		tok = Token{Type: WHITESPACE, Literal: []rune{l.ch}}
		l.isAmongWhiteSpace = false
	case []rune(`ウ`)[0]:
		switch l.peekChar() {
		case []rune(`ホ`)[0]:
			if l.isAmongWhiteSpace {
				tok = take2(X0)
			} else {
				tok = take2(PREFIX)
				l.isAmongWhiteSpace = true
			}
		case []rune(`ウ`)[0]:
			tok = take2(X1)
		case []rune(`ォ`)[0]:
			tok = take2(X6)
		case []rune(`ッ`)[0]:
			c1 := l.ch
			l.readChar()
			c2 := l.ch
			if l.peekChar() == []rune(`ホ`)[0] {
				l.readChar()
				literal := []rune{c1, c2, l.ch}
				tok = Token{Type: XF, Literal: literal}
			} else { // Error
				tok = Token{Type: ILLEGAL, Literal: []rune{c1, c2}}
			}
		default: // Error
			tok = Token{Type: ILLEGAL, Literal: []rune{l.ch}}
		}
	case []rune(`ホ`)[0]:
		switch l.peekChar() {
		case []rune(`ホ`)[0]:
			tok = take2(X2)
		case []rune(`ゥ`)[0]:
			tok = take2(X4)
		case []rune(`ッ`)[0]:
			tok = take2(X5)
		default: // Error
			tok = Token{Type: ILLEGAL, Literal: []rune{l.ch}}
		}
	case []rune(`ゥ`)[0]:
		switch l.peekChar() {
		case []rune(`ホ`)[0]:
			tok = take2(X3)
		case []rune(`ゥ`)[0]:
			tok = take2(XA)
		default: // Error
			tok = Token{Type: ILLEGAL, Literal: []rune{l.ch}}
		}
	case []rune(`う`)[0]:
		switch l.peekChar() {
		case []rune(`ほ`)[0]:
			tok = take2(X7)
		default: // Error
			tok = Token{Type: ILLEGAL, Literal: []rune{l.ch}}
		}
	case []rune(`オ`)[0]:
		switch l.peekChar() {
		case []rune(`ホ`)[0]:
			tok = take2(X9)
		default: // Error
			tok = Token{Type: ILLEGAL, Literal: []rune{l.ch}}
		}
	case []rune(`ッ`)[0]:
		switch l.peekChar() {
		case []rune(`ッ`)[0]:
			tok = take2(XB)
		case []rune(`！`)[0]:
			tok = take2(XC)
		default: // Error
			tok = Token{Type: ILLEGAL, Literal: []rune{l.ch}}
		}
	case []rune(`ォ`)[0]:
		switch l.peekChar() {
		case []rune(`！`)[0]:
			tok = take2(XD)
		default: // Error
			tok = Token{Type: ILLEGAL, Literal: []rune{l.ch}}
		}
	case []rune(`！`)[0]:
		switch l.peekChar() {
		case []rune(`！`)[0]:
			tok = take2(XE)
		default: // Error
			tok = Token{Type: ILLEGAL, Literal: []rune{l.ch}}
		}
	case []rune(`？`)[0]:
		tok = Token{Type: X8, Literal: []rune{l.ch}}
	case 0: // End
		tok = Token{Type: EOF, Literal: []rune{l.ch}}
		l.isAmongWhiteSpace = false
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
