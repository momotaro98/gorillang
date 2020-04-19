package gorillang

import (
	"strconv"
	"strings"
)

type CodePoint string

type Sentence struct {
	CodePoints []CodePoint
}

func (s Sentence) String() string {
	ss := make([]string, len(s.CodePoints))
	for i, cp := range s.CodePoints {
		char, _ := strconv.ParseUint(string(cp), 16, 32)
		r := rune(char)
		ss[i] = string(r)
	}
	return strings.Join(ss, ``)
}
