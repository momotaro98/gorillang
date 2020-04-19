package gorillang

import (
	"fmt"
)

func Decode(input string) string {
	l := NewLexer(input)
	p := NewParser(l)
	sentence := p.ParseSentence()
	return fmt.Sprintf("%s", sentence)
}
