package funcs

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func SplitNl(str string) []string {
	word := ""
	splitedword := []string{}
	skip := false
	for i := 0; i < len(str); i++ {
		if skip {
			skip = false
			continue
		}
		if i != len(str)-1 && str[i] == '\\' && str[i+1] == 'n' {
			if word != "" {
				splitedword = append(splitedword, word)
			}
			word = ""
			skip = true
			splitedword = append(splitedword, "\n")
			continue
		}
		word = word + string(str[i])
	}
	splitedword = append(splitedword, word)
	return splitedword
}

func GetLettres(fileContent []byte) [][]string {
	lettres := [][]string{}
	lettre := []string{}
	line := []byte{}
	s:=""
	for i:=0;i<len(fileContent);i++{
		if fileContent[i]!=13{
			s=s+string(fileContent[i])
		}
	}
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && s[i] == '\n' && s[i+1] == '\n' {
			lettre = append(lettre, string(line))
			lettres = append(lettres, lettre)
			lettre = nil
			line = nil
			continue
		}
		if s[i] == '\n' {
			lettre = append(lettre, string(line))
			line = nil
			continue
		}
		line = append(line, s[i])
	}
	lettres = append(lettres, lettre)
	return lettres
}


func CheckArgs(sstr []string) {
	kinds := []string{"standard", "shadow", "thinkertoy"}
	alignments := []string{"--align=justify", "--align=right", "--align=left", "--align=center"}
	if len(sstr) < 3 {
		log.Fatal("please provide me with 3 arguments")
		return
	}

	for _, arg := range sstr {
		for _, char := range arg {
			if char < 32 || char > 126 {
				log.Fatal("I only accept chars between 32 and 126 on ASCII.")
			}
		}
	}

	if !contains(alignments, sstr[0]) {
		log.Fatal("the alignment should be either one of these: --align=justify, --align=right, --align=left")
	}

	if !contains(kinds, sstr[2]) {
		log.Fatal("the type should be either one of these: standard, shadow, thinkertoy")
	}
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func TerminaleSize() int {
	// Run the "tput cols" command to get the terminal width
	cmd := exec.Command("tput", "cols")
	output, err := cmd.Output()
	if err != nil {
		return 0
	}

	// Convert the output to an integer
	width, err := strconv.Atoi(strings.TrimSpace(string(output)))
	if err != nil {
		return 0
	}

	return width
}
