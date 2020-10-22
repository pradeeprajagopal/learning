package main

import (
	"fmt"
)

//The one with inbuild functions
/*
func numJewelsInStones(J string, S string) int {
	count := 0
	for _, i := range S {
		if strings.ContainsRune(J, i) {
			count++
		}
	}
	return count
}
*/
func numJewelsInStones(J string, S string) int {
	count := 0
	j := make(map[rune]bool, len(J))
	for _, v := range J {
		j[v] = true
	}
	for _, i := range S {
		if j[i] {
			count++
		}
	}
	return count
}
func main() {
	fmt.Println(numJewelsInStones("z", "ZZ"))
}
