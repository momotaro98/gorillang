package gorillang

import (
	"strconv"
	"strings"
)

var encodeMap = map[string]TokenType{
	"0": X0,
	"1": X1,
	"2": X2,
	"3": X3,
	"4": X4,
	"5": X5,
	"6": X6,
	"7": X7,
	"8": X8,
	"9": X9,
	"a": XA,
	"b": XB,
	"c": XC,
	"d": XD,
	"e": XE,
	"f": XF,
}

func Encode(input string) string {
	codePoints := subEncode(input)

	encodedSlice := make([]string, len(codePoints))
	for i, cp := range codePoints {
		var encoded string
		encoded += PREFIX // 文字の頭に"ウホ"(PREFIX)が必要
		for j := range cp {
			encoded += string(encodeMap[string(cp[j])])
		}
		encodedSlice[i] = encoded
	}

	return strings.Join(encodedSlice, WHITESPACE)
}

func subEncode(input string) []CodePoint {
	rs := []rune(input)
	cps := make([]CodePoint, len(rs))
	for i, r := range rs {
		v := int64(r)
		s16 := strconv.FormatInt(v, 16)
		cps[i] = CodePoint(s16)
	}
	return cps
}
