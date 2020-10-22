package main

func isMonotonic(A []int) bool {
	inc, dec := true, true
	for i := 1; i < len(A); i++ {
		inc, dec = inc && (A[i-1] <= A[i]), dec && (A[i-1] >= A[i])
	}
	return inc || dec
}
