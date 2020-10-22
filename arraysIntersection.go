package main

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	var result []int
	var a1, a2 = make(map[int]bool, len(arr1)), make(map[int]bool, len(arr2))
	for _, i := range arr1 {
		a1[i] = true
	}
	for _, j := range arr2 {
		a2[j] = true
	}
	for _, k := range arr3 {
		if a1[k] && a2[k] {
			result = append(result, k)
		}
	}
	return result
}
