package main

func containsDuplicate(nums []int) bool {
	v := make(map[int]bool)
	for _, i := range nums {
		if v[i] {
			return true
		} else {
			v[i] = true
		}
	}
	return false
}
