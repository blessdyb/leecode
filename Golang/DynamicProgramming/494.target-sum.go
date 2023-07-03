package main

func findTargetSumWaysRecursive(nums []int, target int) int {
	ret := 0
	var recursive func(nums []int, t int)
	recursive = func(nums []int, t int) {
		if len(nums) == 0 {
			if t == 0 {
				ret++
			}
			return
		}
		recursive(nums[1:], t+nums[0])
		recursive(nums[1:], t-nums[0])
	}
	recursive(nums, target)
	return ret
}

func findTargetSumWays2DDPI(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2*sum+1)
	}
	// dp[0][0] = 1, without offset sum
	dp[0][nums[0]+sum] = 1
	// consider nums[0] could be 0
	dp[0][-nums[0]+sum] += 1
	for i := 1; i < len(nums); i++ {
		for j := -sum; j <= sum; j++ {
			if dp[i-1][j+sum] > 0 {
				dp[i][j+nums[i]+sum] += dp[i-1][j+sum]
				dp[i][j-nums[i]+sum] += dp[i-1][j+sum]
			}
		}
	}
	if target > sum || target < -sum {
		return 0
	}
	return dp[len(nums)-1][target+sum]
}

/ **
 * Let p, n be the sums of numbers with positive and negative symbols, respectively. 
 * Let sum = p + n, target = p - n
 * So p = (sum + target) / 2, n = (sum - target) / 2
 * It means we need to find integers p or n
 * We successfully converted this problem to a 0-1 backpack problem
 * /
func findTargetSumWays2DDPII(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if target > sum || target < -sum || (target + sum) % 2 == 1 {
		return 0
	}
	size := (target + sum) / 2
	dp := make([][]int, len(nums) + 1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]int, size + 1)
	}
	dp[0][0] = 1
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= size; j++ {
			dp[i][j] = dp[i - 1][j]
			if j >= nums[i - 1] {
				dp[i][j] += dp[i - 1][j - nums[i - 1]]
			}
		}
	}
	return dp[len(nums)][size]
}

func findTargetSumWays1DDPII(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if target > sum || target < -sum || (sum + target) % 2 == 1 {
		return 0
	}
	size := (sum + target) / 2
	dp := make([]int, size + 1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := size; j >= nums[i]; j-- {
			dp[j] += dp[j - nums[i]]
		}
	}
	return dp[size]
}