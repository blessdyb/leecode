package main

// If x ^ y = z, then x ^ z = y and x = y ^ z
// Since k ~ [0, 2^maximumBit - 1]
// Given the query  nums[0]^nums[1]^nums[2]^...^nums[len(nums)-1]^k = 2^maximumBit - 1
// So k = nums[0]^nums[1]^nums[2]^...^nums[len(nums)-1]^(2^maximumBit-1)
func getMaximumXor(nums []int, maximumBit int) []int {
	all := 0
	for _, num := range nums {
		all ^= num
	}
	max := (2 << (maximumBit - 1)) - 1
	ret := []int{}
	for len(nums) > 0 {
		ret = append(ret, all^max)
		all ^= nums[len(nums)-1]
		nums = nums[:len(nums)-1]
	}
	return ret
}

func getMaximumXorTLE(nums []int, maximumBit int) []int {
	length := len(nums)
	preXor := make([]int, length)
	preXor[0] = nums[0]
	for i := 1; i < length; i++ {
		preXor[i] = preXor[i-1] ^ nums[i]
	}
	ret := []int{}
	k := 2 << (maximumBit - 1)
	for i := 1; i <= length; i++ {
		max := 0
		index := 0
		for j := 0; j < k; j++ {
			temp := preXor[len(nums)-1] ^ j
			if max < temp {
				max = temp
				index = j
			}
		}
		ret = append(ret, index)
		nums = nums[:len(nums)-1]
	}
	return ret
}
