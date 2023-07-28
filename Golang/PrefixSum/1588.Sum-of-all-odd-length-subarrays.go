package main

func sumOddLengthSubarrays(arr []int) int {
	sum := make([]int, len(arr)+1)
	sum[0] = 0
	for i := 1; i <= len(arr); i++ {
		sum[i] = sum[i-1] + arr[i-1]
	}
	ret := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j <= len(arr); j++ {
			if (j-i)%2 == 1 {
				ret += sum[j] - sum[i]
			}
		}
	}
	return ret
}
