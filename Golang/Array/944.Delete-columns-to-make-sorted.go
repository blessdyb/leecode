package main

func minDeletionSize(strs []string) int {
	ret := 0
	for j := 0; j < len(strs[0]); j++ {
		for i := 1; i < len(strs); i++ {
			if strs[i][j] < strs[i-1][j] {
				ret++
				break
			}
		}
	}
	return ret
}
