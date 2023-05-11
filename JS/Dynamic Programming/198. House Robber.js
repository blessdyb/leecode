/**
 * @param {number[]} nums
 * @return {number}
 */
var rob = function(nums) {
    const dp = [0, nums[0]];
    const n = nums.length;
    for(let i = 2; i <= n; i++) {
        dp[i] = Math.max(nums[i - 1] + dp[i - 2], dp[i - 1]);
    }
    return dp[n];
};
