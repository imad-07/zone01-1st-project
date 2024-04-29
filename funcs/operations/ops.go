package operations

import (
	"strconv"
	"strings"
)

func Cap(s *string) {
	rslt := []rune(*s)
	if len(rslt) > 0 && rslt[0] >= 'a' && rslt[0] <= 'z' {
		rslt[0] -= 32
	}
	for i := 1; i < len(rslt); i++ {
		if rslt[i] >= 'a' && rslt[i] <= 'z' {
			rslt[i] -= 32
		}
		if rslt[i] >= 'A' && rslt[i] <= 'Z' {
			rslt[i] += 32
		}
	}
	*s = string(rslt)
}

func ToLower(s *string) {
	x := []rune(*s)
	for i := 0; i < len(x); i++ {
		if x[i] >= 'A' && x[i] <= 'Z' {
			x[i] += 32
		}
	}
	*s = string(x)
}

func ToUpper(s *string) {
	x := []rune(*s)
	for i := 0; i < len(x); i++ {
		if x[i] >= 'a' && x[i] <= 'z' {
			x[i] -= 32
		}
	}
	*s = string(x)
}

func ToBin(s *string) {
	base, _ := strconv.ParseInt((*s), 2, 64)
	(*s) = strconv.FormatInt(base, 10)
}

func ToHex(s *string) {
	base, _ := strconv.ParseInt((*s), 16, 64)
	(*s) = strconv.FormatInt(base, 10)
}

func ClrIndex(input *[]string, index int) {
	*input = append((*input)[:index], (*input)[index+1:]...)
}

func Clean(input *string) {
	words := strings.Fields(*input)
	*input = strings.Join(words, " ")
}
