package main

import (
	"fmt"
	"strings"
)

func restoreString(s string, indices []int) string {
	var result = make([]string, len(indices))
	for i, S := range s {
		result[indices[i]] = string(S)
	}
	return strings.Join(result, "")
}

func main() {
	i := []int{4, 5, 6, 7, 0, 2, 1, 3}
	fmt.Println(restoreString("codeleet", i))
}
