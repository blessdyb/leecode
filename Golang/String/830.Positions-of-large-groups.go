package main

func largeGroupPositions(s string) [][]int {
	ret := [][]int{}
	temp := []int{0}
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			temp = append(temp, i-1)
			if temp[1]-temp[0]+1 >= 3 {
				ret = append(ret, temp)
			}
			temp = []int{i}
		}
	}
	// take care of the last index
	temp = append(temp, len(s)-1)
	if temp[1]-temp[0]+1 >= 3 {
		ret = append(ret, temp)
	}
	return ret
}
