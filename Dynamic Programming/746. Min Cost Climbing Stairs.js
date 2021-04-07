/**
 * @param {number[]} cost
 * @return {number}
 */
var minCostClimbingStairs = function(cost) { // DP
    const dp = [0, 0]; // Since you can only take one / two steps, so dp[1] = 0, dp[2] = Math.min(cost[0], cost[1]));
    for (let i = 2; i <= cost.length; i++) {
        dp[i] = Math.min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2]);
    }
    return dp[cost.length];
};
