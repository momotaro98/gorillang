package gorillang

type Token struct {
	Type    TokenType
	Literal []rune
}

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	WHITESPACE = ` `

	PREFIX = "ウホ"

	X0 = "ウホ"
	X1 = "ウウ"
	X2 = "ホホ"
	X3 = "ゥホ"
	X4 = "ホゥ"
	X5 = "ホッ"
	X6 = "ウォ"
	X7 = "うほ"
	X8 = "？"
	X9 = "オホ"
	XA = "ゥゥ"
	XB = "ッッ"
	XC = "ッ！"
	XD = "ォ！"
	XE = "！！"
	XF = "ウッホ"
)