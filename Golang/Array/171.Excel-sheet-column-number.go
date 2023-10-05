package main

func titleToNumber(columnTitle string) int {
	ret := 0
	for i := 0; i < len(columnTitle); i++ {
		ret = ret*26 + int(columnTitle[i]-'A'+1)
	}
	return ret
}
