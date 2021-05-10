/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var findTargetSumWays = function(nums, target) {
    let result = 0;
    const n = nums.length;
    function helper(i, t) {
        if (i === 0) {
            if (nums[0] === t) { // In case nums[i] is 0, we need to count it twice
                result++;
            }
            if(nums[0] === -t) {
                result++;
            }
        } else {
            helper(i - 1, t + nums[i]);
            helper(i - 1, t - nums[i]);
        }
    }
    helper(n - 1, target);
    return result;
};

var findTargetSumWays = function(nums, target) {
    const n = nums.length;
    const sum = nums.reduce((acc, item) => acc + item, 0);
    const dp = new Array(n + 1).fill(0).map(item => new Array(2 * sum + 1).fill(0));
    dp[0][sum] = 1; // given target 0, with 0 nums, we have 1 way
    for(let i = 1; i <= n; i++) {
        for(let j = -sum; j <= sum; j++) {
            if (j + nums[i - 1] <= sum) {
                dp[i][j + sum + nums[i - 1]] += dp[i - 1][j + sum];   
            }
            if (j - nums[i - 1] >= -sum) {
                dp[i][j + sum - nums[i - 1]] += dp[i - 1][j + sum];   
            }
        }
    }
    return target > sum || target < -sum ? 0 : dp[n][target + sum];
};
