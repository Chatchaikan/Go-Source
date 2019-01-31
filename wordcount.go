package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	strs := strings.Fields(s)
	res := make(map[string]int)

	for _, str := range strs {
		res[strings.ToLower(str)]++
	}
	return res
	//return map[string]int{"x": 1}
}

func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	w := WordCount(s) // HL
	fmt.Println(w)

//func main() {
//	fmt.Println(WordCount("If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."))
}