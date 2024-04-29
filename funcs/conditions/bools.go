package conditions

import "operations/funcs/operations"

func IsAlpha(s string) bool {
	ss := []rune(s)
	for i := 0; i < len(ss); i++ {
		if (ss[i] >= 'a' && ss[i] <= 'z') || (ss[i] >= 'A' && ss[i] <= 'Z') {
			return true
		}
	}
	return false
}

func IsVowel(str string) bool {
	operations.ToLower(&str)
	s := []rune(str)
	return s[0] == 'a' || s[0] == 'e' || s[0] == 'i' || s[0] == 'o' || s[0] == 'u' || s[0] == 'h'
}

func IsPonc(s string) bool {
	for _, char := range s {
		if char != '.' && char != ',' && char != '!' && char != '?' && char != ';' && char != ':' {
			return false
		}
	}
	return true
}

func IsNumerik(str string) bool {
	ss := []rune(str)
	for i := 0; i < len(ss); i++ {
		if ss[i] < '0' || ss[i] > '9' {
			return false
		}
	}
	return true
}

func IsAlphaNumerik(str string) bool {
	return IsAlphaNumerik(str) || IsAlpha(str)
}
