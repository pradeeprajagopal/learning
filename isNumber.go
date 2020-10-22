package main

import (
	"fmt"
	"strings"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	founde := false
	isSign := false
	dotfound := false
	if len(s) == 0 || s[0] == 'e' || (s[0] == '.' && len(s) == 1) || s[len(s)-1] == '+' || s[len(s)-1] == '-' {
		return false
	}

	for i := range s {
		if !((s[i] <= '9' && s[i] >= '0') || (s[i] == '+') || (s[i] == '-') || (s[i] == 'e') || (s[i] == '.') || s[i] == ' ') {
			return false
		}
		if s[i] == '.' {
			if dotfound {
				return false
			}
			if (i < len(s)-1) && s[i+1] == 'e' {
				return false
			}
			dotfound = true
		}
		if founde && s[i] == '.' {
			return false
		}

		if (s[i] == '+') || (s[i] == '-') {
			if isSign || (!founde && i != 0) {
				return false
			}
			isSign = true
		}
		if s[i] == 'e' {
			if !founde {
				if i == len(s)-1 {
					return false
				}
				founde = true

			} else {
				return false
			}
		}
	}
	return true

}

// func isNumber(s string) bool {
// 	s = strings.TrimSpace(s)
// 	length := len(s)
// 	initial := 0
// 	if length == 0 {
// 		return false
// 	}
// 	for initial < length {
// 		if s[initial] == '+' || s[initial] == '-' {
// 			initial++
// 			break
// 		} else if s[initial] <= '9' && s[initial] >= '0' {
// 			break
// 		}
// 		return false
// 	}
// 	fmt.Println(initial)
// 	fmt.Println(length)

// 	return true
// }
func main() {
	fmt.Println(isNumber("0"))         //met
	fmt.Println(isNumber(" 0.1 "))     //met
	fmt.Println(isNumber("abc"))       //met
	fmt.Println(isNumber("1 a"))       //met
	fmt.Println(isNumber("2e10"))      //met
	fmt.Println(isNumber(" -90e3   ")) //met
	fmt.Println(isNumber(" 1e"))       //met
	fmt.Println(isNumber("e3"))        //met
	fmt.Println(isNumber(" 6e-1"))     //met
	fmt.Println(isNumber(" 99e2.5 "))  //met
	fmt.Println(isNumber("53.5e93"))   //met
	fmt.Println(isNumber(" --6 "))     //met
	fmt.Println(isNumber("-+3"))       //met
	fmt.Println(isNumber("95a54e53"))  //met
	fmt.Println(isNumber("6+1"))       //met
	fmt.Println(isNumber(".e1"))       //met
	fmt.Println(isNumber("4e+"))       //met
}
