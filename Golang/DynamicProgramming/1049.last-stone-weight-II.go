package main

// 0-1 backpack, given target volum sum / 2 and the stone i with weight stones[i] and value[i], get the largest value
// then sum - dp[target] - dp[target] is the result
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum / 2
	dp := make([]int, target+1)
	dp[0] = 0
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			if dp[j] < dp[j-stones[i]]+stones[i] {
				dp[j] = dp[j-stones[i]] + stones[i]
			}
		}
	}
	return sum - dp[target] - dp[target]
}
