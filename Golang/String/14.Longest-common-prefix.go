package main

func longestCommonPrefix(strs []string) string {
	queue := []byte(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(queue) == 0 {
			break
		} else if len(queue) > len(strs[i]) {
			queue = queue[:len(strs[i])]
		}
		for j := 0; j < len(strs[i]) && j < len(queue); j++ {
			if queue[j] != strs[i][j] {
				queue = queue[:j]
				break
			}
		}
	}
	return string(queue)
}
