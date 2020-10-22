package main

import (
	"fmt"
	"strings"
)

var m = map[int]string{
	0:          "Zero",
	1:          "One",
	2:          "Two",
	3:          "Three",
	4:          "Four",
	5:          "Five",
	6:          "Six",
	7:          "Seven",
	8:          "Eight",
	9:          "Nine",
	10:         "Ten",
	11:         "Eleven",
	12:         "Twelve",
	13:         "Thirteen",
	14:         "Fourteen",
	15:         "Fifteen",
	16:         "Sixteen",
	17:         "Seventeen",
	18:         "Eighteen",
	19:         "Nineteen",
	20:         "Twenty",
	30:         "Thirty",
	40:         "Forty",
	50:         "Fifty",
	60:         "Sixty",
	70:         "Seventy",
	80:         "Eighty",
	90:         "Ninety",
	100:        "Hundred",
	1000:       "Thousand",
	1000000:    "Million",
	1000000000: "Billion",
}

func numberToWords(num int) string {
	var result []string
	if num == 0 {
		return m[0]
	}
	for _, order := range []int{1000000000, 1000000, 1000, 1} {
		v := num / order
		if v > 0 {
			num -= v * order
			if v >= 100 {
				result = append(result, m[v/100])
				result = append(result, m[100])
				v -= (v / 100) * 100
			}
			if v > 20 {
				result = append(result, m[(v/10)*10])
				v -= (v / 10) * 10
			}
			if v > 0 {
				result = append(result, m[v])
				v -= v
			}
			if order > 100 {
				result = append(result, m[order])
			}
		}

	}
	return strings.Join(result, " ")
}

func main() {
	fmt.Println(numberToWords(4560000001))
}
