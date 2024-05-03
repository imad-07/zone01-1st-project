package main

import (
	"fmt"
	"os"

	"funcs/funcs"
)

func main() {
	args := os.Args[1:]
	funcs.CheckArgs(args)
	//alignement := args[0]
	alignement := args[0][8:]
	victim := args[1]
	kind := args[2] + ".txt"
	x := funcs.TerminaleSize()
	fmt.Println(alignement, victim, kind, x)
	// here we split our input with new lines while keeping each one of them in an indexed place in the array
	word := funcs.SplitNl(victim)
	fileContent, err := os.ReadFile(kind)
	if err != nil {
		fmt.Printf("error in the kind file")
		return
	}
	// here we get the standard art from the file that they gave us
	lettres := funcs.GetLettres(fileContent)
	bl := false
	for l := 0; l < len(word); l++ {
		if word[l] == "" {
			continue
		}
		if word[l] == "\n" {
			if l == len(word)-1 {
				continue
			}
			if bl && word[l+1] != "\n" {
				continue
			}
			fmt.Printf("\n")
			continue
		}
		for i := 1; i < 9; i++ {
			for j := 0; j < len(word[l]); j++ {
				fmt.Printf(lettres[word[l][j]-32][i])
			}
			fmt.Print("\n")
		}
		bl = true
	}
}
