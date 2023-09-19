package main

func letterCasePermutation(s string) []string {
	hashmap := map[string]bool{}
	permutation := []byte{}
	var backtracking func(int)
	backtracking = func(i int) {
		if len(permutation) == len(s) {
			hashmap[string(permutation)] = true
		} else {
			for j := i; j < len(s); j++ {
				c := s[j]
				permutation = append(permutation, c)
				backtracking(j + 1)
				permutation = permutation[:len(permutation)-1]
				if c >= '0' && c <= '9' {
					continue
				} else if c >= 'a' && c <= 'z' {
					c -= 'a' - 'A'
				} else {
					c += 'a' - 'A'
				}
				permutation = append(permutation, c)
				backtracking(j + 1)
				permutation = permutation[:len(permutation)-1]
			}
		}
	}
	backtracking(0)
	ret := []string{}
	for k := range hashmap {
		ret = append(ret, k)
	}
	return ret
}
