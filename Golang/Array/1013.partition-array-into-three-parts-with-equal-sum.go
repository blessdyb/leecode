package main

func canThreePartsEqualSum(arr []int) bool {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	if sum%3 != 0 {
		return false
	}
	count, temp, sample := 0, 0, sum/3
	for i := 0; i < len(arr); i++ {
		temp += arr[i]
		if temp == sample {
			count++
			temp = 0
		}
	}
	// Since sum % 3 == 0, count >= 3 will work
	return count >= 3
}
