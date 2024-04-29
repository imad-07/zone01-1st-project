package punctuations

import (
	"regexp"

	"operations/funcs/conditions"
)

func Anoa(s *[]string) {
	for i := 0; i < len(*s)-1; i++ {
		if (string((*s)[i]) == "a" || string((*s)[i]) == "A") && conditions.IsVowel(string((*s)[i+1][0])) && string((*s)[i+1][0]) != "" {
			(*s)[i] += "n"
		}
	}
}

func Poncword(s *[]string) {
	filler := ""
	for i := 1; i < len(*s); i++ {
		if (*s)[i] != "" && conditions.IsPonc(string((*s)[i][0])) {
			for j := 0; j < len((*s)[i])-1; j++ {
				if conditions.IsPonc(string((*s)[i][j])) {
					filler += string((*s)[i][j])
				}
				if !conditions.IsPonc(string((*s)[i][j+1])) {
					break
				}
			}
			(*s)[i-1] += filler
			(*s)[i] = (*s)[i][len(filler):]
		}
		filler = ""
	}
}

func Poncs(s *string) {
	x := []rune(*s)
	for i := 1; i < len(x); i++ {
		if conditions.IsPonc(string(x[i])) && (i == 1 || x[i-1] == ' ') {
			wordStart := i
			for wordStart > 0 && x[wordStart-1] == ' ' {
				wordStart--
			}
			for j := wordStart; j < i; j++ {
				x[j], x[i] = x[i], x[j]
				i--
			}
			if i+1 < len(x) {
				x[i+1] = ' '
			}
		}
	}
	for i := 1; i < len(x)-1; i++ {
		if conditions.IsPonc(string(x[i])) && conditions.IsAlpha(string(x[i+1])) {
			for k := len(x) - 1; k > i+1; k-- {
				x[k] = x[k-1]
			}
			x[i+1] = ' ' 
			i++         
		}
	}
	*s = string(x)
}

func Sinq(str *string) {
	text := *str
	s := []string{}
	j := 0
	c := false

	for i := 0; i < len(text); i++ {
		if text[i] == '\'' {
			if i != 0 && i != len(text)-1 {
				if text[i-1] == ' ' || text[i+1] == ' ' {
					if !c {
						s = append(s, text[j:i])
						j = i
						c = true
					} else {
						s = append(s, text[j:i+1])
						j = i + 1
						c = false
					}
				}
			} else {
				if !c {
					s = append(s, text[j:i])
					j = i
					c = true
				} else {
					s = append(s, text[j:i+1])
					j = i + 1
					c = false
				}
			}
		}
	}
	s = append(s, text[j:])

	t := ""
	for i := 0; i < len(s); i += 1 {
		re := regexp.MustCompile(`' *(.*[\S]) *'`)
		s[i] = re.ReplaceAllString(s[i], "'$1'")
		t += s[i]
	}

	*str = t
}
