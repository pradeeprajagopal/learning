package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	shash := make(map[rune]int)
	for _, l := range s {
		shash[l]++
	}
	for _, j := range t {
		if shash[j] == 0 {
			return false
		}
		shash[j]--
	}
	return true

}
