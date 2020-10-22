package main

import (
	"fmt"
	"strconv"
)

/*
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"


Constraints:

Each string consists only of '0' or '1' characters.
1 <= a.length, b.length <= 10^4
Each string is either "0" or doesn't contain any leading zero.
*/
func addBinary(a string, b string) string {
	//Initialize i and j for the iteration. Result as empty string and carryfwd as 0 to start with
	i, j, res, carryfwd := len(a)-1, len(b)-1, "", 0
	//for to make sure to cover all character is covered we are going on a loop until both input parameters are empty
	for i >= 0 || j >= 0 {
		//adding the carryfwd to append the value
		sum := carryfwd
		//till i is not equal to zero
		if i >= 0 {
			//append the value to sum
			sum += int(a[i] - '0')
			//decremenet the i
			i--
		}
		//till i is not equal to zero
		if j >= 0 {
			//append the value to sum
			sum += int(b[j] - '0')
			//decremenet the j
			j--
		}
		//convert int to string plus append the previous value at the end
		res = strconv.Itoa(sum%2) + res
		//carryfwd is added to the next
		carryfwd = sum / 2
	}
	//after i and j are zero need to validate if carryfwd exist then append it at front
	if carryfwd == 1 {
		res = "1" + res
	}
	return res
}

// func addBinaryRecursion(a string, b string) string {
// 	i, j, res, carryfwd := len(a)-1, len(b)-1, "", 0
// 	for i >= 0 || j >= 0 {
// 		if i != 0 && j != 0 {
// 			i--
// 			j--
// 		}
// 		if i != 0 && j == 0 {
// 			i--
// 		}
// 		if j != 0 && i == 0 {
// 			j--
// 		}
// 	}

// }

func additem(a int, b int) (int, int) {
	if a == 1 && b == 1 {
		return 0, 1
	} else if (a == 1 || b == 1) && (a == 0 || b == 0) {
		return 1, 0
	} else if a == 0 || b == 0 {
		return 0, 0
	}
	return 0, 0
}
func main() {
	fmt.Println(addBinary("11", "1"))
}
