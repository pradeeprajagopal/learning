package main

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	l := make([]int, len(nums))
	r := make([]int, len(nums))
	l[0], r[len(nums)-1] = 1, 1
	for i, j := 0, len(nums)-1; i < len(nums) && j >= 0; i, j = i+1, j-1 {
		l[i] = l[i-1] * nums[i-1]
		r[j] = r[j+1] * nums[j+1]
	}
	for i := range ans {
		ans[i] = l[i] * r[i]
	}
	return ans
}
