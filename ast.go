package gorillang

import (
	"strconv"
	"strings"
)

type CPElement string

type Sentence struct {
	CodePoint []CPElement
}

func (s Sentence) String() string {
	ss := make([]string, len(s.CodePoint))
	for i, cp := range s.CodePoint {
		char, _ := strconv.ParseUint(string(cp), 16, 32)
		r := rune(char)
		ss[i] = string(r)
	}
	return strings.Join(ss, ``)
}
