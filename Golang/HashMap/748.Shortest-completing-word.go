package main

func shortestCompletingWord(licensePlate string, words []string) string {
	hashmap := map[byte]int{}
	var ret string
	for i := 0; i < len(licensePlate); i++ {
		if licensePlate[i] >= 'a' && licensePlate[i] <= 'z' {
			hashmap[licensePlate[i]]++
		} else if licensePlate[i] >= 'A' && licensePlate[i] <= 'Z' {
			hashmap[licensePlate[i]+'a'-'A']++
		}
	}
	for _, word := range words {
		flag := true
		temp := map[byte]int{}
		for i := 0; i < len(word); i++ {
			temp[word[i]]++
		}
		for k, v := range hashmap {
			if temp[k] < v {
				flag = false
				break
			}
		}
		if flag && (ret == "" || len(ret) > len(word)) {
			ret = word
		}
	}
	return ret
}
