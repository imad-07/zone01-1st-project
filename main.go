package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"operations/funcs/conditions"
	"operations/funcs/operations"
	"operations/funcs/punctuations"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("error: wrp)ong input")
	} else {
		stxt := args[0]
		etxt := args[1]
		content, err := os.ReadFile(stxt)
		if err != nil {
			log.Fatal(err)
		}
		var beta []string
		content2 := strings.Split(string(content), "\n")
		for i := 0; i < len(content2); i++ {
			content2[i] = strings.TrimRight(content2[i], " ")
		}
		for i := 0; i < len(content2); i++ {
			sample := strings.Split(string(content2[i]), " ")
			edit(&sample)
			sample2 := strings.Join(sample, " ")
			punctuations.Poncs(&sample2)
			operations.Clean(&sample2)
			punctuations.Sinq(&sample2)
			beta = append(beta, sample2)
			if i != len(content2)-1 {
				beta = append(beta, "\n")
			}
		}
		result := ""
		for i := 0; i < len(beta); i++ {
			result += beta[i]
		}
		err = os.WriteFile(etxt, []byte(result), 0o644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func edit(victim *[]string) {
	victim0 := (*victim)
	if victim0[0] == "(hex)" || victim0[0] == "(bin)" || victim0[0] == "(up)" || victim0[0] == "(cap)" || victim0[0] == "(low)" {
		operations.ClrIndex(&victim0, 0)
	} else if (victim0[0] == "(up," || victim0[0] == "(low," || victim0[0] == "(cap,")  && len(victim0) > 1{
		if conditions.IsNumerik(victim0[1][:len(victim0[1])-1]) && string(victim0[1][len(victim0[1])-1]) == ")" {
			operations.ClrIndex(&victim0, 0)
			operations.ClrIndex(&victim0, 0)
		}
	}
	for i := len(victim0) - 1; i >= 1; i-- {
		if victim0[i] == "(hex)" {
			operations.ToHex(&victim0[i-1])
			operations.ClrIndex(&victim0, i)
		} else if victim0[i] == "(bin)" {
			operations.ToBin(&victim0[i-1])
			operations.ClrIndex(&victim0, i)
		} else if victim0[i] == "(up)" {
			operations.ToUpper(&victim0[i-1])
			operations.ClrIndex(&victim0, i)
		} else if victim0[i] == "(low)" {
			operations.ToLower(&victim0[i-1])
			operations.ClrIndex(&victim0, i)
		} else if victim0[i] == "(cap)" {
			operations.Cap(&victim0[i-1])
			operations.ClrIndex(&victim0, i)
		} else if victim0[i] == "(up," && i+1 < len(victim0) && string(victim0[i+1][len(victim0[i+1])-1]) == ")" {
			if conditions.IsNumerik(victim0[i+1][:len(victim0[i+1])-1]) {
				x, _ := strconv.Atoi(victim0[i+1][:len(victim0[i+1])-1])
				j := 1
				for j <= x && i-j >= 0 {
					operations.ToUpper(&victim0[i-j])
					if victim0[i-j] != " " {
						j++
					}
				}
				operations.ClrIndex(&victim0, i)
				operations.ClrIndex(&victim0, i)
			}
		} else if victim0[i] == "(low," && i+1 < len(victim0) && string(victim0[i+1][len(victim0[i+1])-1]) == ")"{
			if conditions.IsNumerik(victim0[i+1][:len(victim0[i+1])-1]) {
				x, _ := strconv.Atoi(victim0[i+1][:len(victim0[i+1])-1])
				j := 1
				for j <= x && i-j >= 0 {
					operations.ToLower(&victim0[i-j])
					if victim0[i-j] != " " {
						j++
					}
				}
				operations.ClrIndex(&victim0, i)
				operations.ClrIndex(&victim0, i)
			}
		} else if victim0[i] == "(cap," && i+1 < len(victim0) && string(victim0[i+1][len(victim0[i+1])-1]) == ")"{
			if conditions.IsNumerik(victim0[i+1][:len(victim0[i+1])-1]) {
				x, _ := strconv.Atoi(victim0[i+1][:len(victim0[i+1])-1])
				j := 1
				for j <= x && i-j >= 0 {
					operations.Cap(&victim0[i-j])
					if victim0[i-j] != " " {
						j++
					}
				}
				operations.ClrIndex(&victim0, i)
				operations.ClrIndex(&victim0, i)
			}
		}
	}
	punctuations.Anoa(&victim0)
	*victim = victim0
}
