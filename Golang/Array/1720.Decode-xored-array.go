package main

func decode(encoded []int, first int) []int {
	ret := make([]int, len(encoded)+1)
	ret[0] = first
	for i := 1; i <= len(encoded); i++ {
		ret[i] = ret[i-1] ^ encoded[i-1]
	}
	return ret
}
