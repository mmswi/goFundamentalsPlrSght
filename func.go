package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "function basics"
	author := "mihai marr"

	fmt.Println(converter(module, author))

//	creating a function with multiple arguments
	bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 16, 8)
	fmt.Println(bestFinish)
}

func converter(mod, aut string) (string, string) {
	mod = strings.Title(mod)
	aut = strings.ToUpper(aut)
	return mod, aut
}

// declaring multiple arguments
func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]
	for _,i := range finishes {
		if i<best {
			best = i
		}
	}
	return best
}
