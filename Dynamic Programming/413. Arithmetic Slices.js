/**
 * @param {number[]} nums
 * @return {number}
 */
var numberOfArithmeticSlices = function(nums) {
    const n = nums.length;
    const dp = new Array(n + 1).fill(0);
    for(let i = 2; i <= n; i++) {
        if (nums[i] - nums[i - 1] === nums[i - 1] - nums[i - 2]) {
            dp[i + 1] = dp[i] + 1; 
        }
    }
    return dp.reduce((acc, i) => acc + i, 0);
}
