package main

func mostCommonWord(paragraph string, banned []string) string {
	ret := ""
	max := 0
	hashmap := map[string]int{}
	bannedMap := map[string]bool{}
	for _, b := range banned {
		bannedMap[b] = true
	}
	letters := []byte{}
	for i := 0; i < len(paragraph); i++ {
		if paragraph[i] >= 'a' && paragraph[i] <= 'z' {
			letters = append(letters, paragraph[i])
		} else if paragraph[i] >= 'A' && paragraph[i] <= 'Z' {
			letters = append(letters, paragraph[i]+'a'-'A')
		} else if len(letters) > 0 {
			temp := string(letters)
			if !bannedMap[temp] {
				hashmap[temp]++
				if hashmap[temp] > max {
					ret = temp
					max = hashmap[temp]
				}
			}
			letters = []byte{}
		}
	}
	if len(letters) > 0 {
		temp := string(letters)
		if !bannedMap[temp] {
			hashmap[temp]++
			if hashmap[temp] > max {
				ret = temp
				max = hashmap[temp]
			}
		}
	}
	return ret
}
