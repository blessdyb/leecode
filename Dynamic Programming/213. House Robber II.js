/**
 * @param {number[]} nums
 * @return {number}
 */
var rob = function(nums) {
    function helper(nums) {
        const dp = [0, nums[0]];
        const n = nums.length;
        for(let i = 2; i <= n; i++) {
            dp[i] = Math.max(nums[i - 1] + dp[i - 2], dp[i - 1]);
        }
        return dp[n];
    }
    const n = nums.length;
    return Math.max(nums[0] + helper(nums.slice(2, n - 1)), helper(nums.slice(1)))
};
