package gorillang

var codePointMap = map[TokenType]CodePoint{
	X0: "0",
	X1: "1",
	X2: "2",
	X3: "3",
	X4: "4",
	X5: "5",
	X6: "6",
	X7: "7",
	X8: "8",
	X9: "9",
	XA: "A",
	XB: "B",
	XC: "C",
	XD: "D",
	XE: "E",
	XF: "F",
}

type Parser struct {
	l *Lexer

	curToken  Token
	peekToken Token
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{l: l}

	// 2つトークンを読み込む。curToken と peekToken の両方がセットされる。
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseSentence() *Sentence {
	sentence := &Sentence{}
	sentence.CodePoints = []CodePoint{}

	for p.curToken.Type != EOF {
		cp := p.parseSentence()
		if cp != "" {
			sentence.CodePoints = append(sentence.CodePoints, cp)
		}
		p.nextToken()
	}

	return sentence
}

func (p *Parser) parseSentence() CodePoint {
	var codePoint CodePoint

	if p.curToken.Type != PREFIX {
		return "" // TODO: error
	}

	p.nextToken()

	return p.recursiveFn(codePoint)
}

func (p *Parser) recursiveFn(codePoint CodePoint) CodePoint {
	switch tType := p.curToken.Type; tType {
	case ILLEGAL:
		return codePoint // TODO: error
	case EOF, WHITESPACE:
		return codePoint
	default:
		codePoint += codePointMap[tType]
		p.nextToken()
		return p.recursiveFn(codePoint)
	}
}
