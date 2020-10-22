package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	result := strings.ReplaceAll(address, ".", "[.]")
	return result
}

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}
