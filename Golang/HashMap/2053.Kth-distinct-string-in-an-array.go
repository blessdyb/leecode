package main

func kthDistinct(arr []string, k int) string {
	hashmap := map[string]int{}
	for _, str := range arr {
		hashmap[str]++
	}
	ret := ""
	for i := 0; i < len(arr); i++ {
		if hashmap[arr[i]] == 1 {
			k--
			if k == 0 {
				ret = arr[i]
				break
			}
		}
	}
	return ret
}
