package main

func rotatedDigits(N int) int {
	count := 0
	for i := 1; i <= N; i++ {
		if isValid(i) {
			count++
		}
	}
	return count
}

func isValid(n int) bool {
	result := false
	for ; n > 0; n /= 10 {
		switch n % 10 {
		case 2, 5, 6, 9:
			result = true
		case 3, 4, 7:
			return false
		}
	}
	return result
}
