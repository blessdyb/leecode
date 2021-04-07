/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function(n) {
    const dp = [0, 1, 2];
    for(let i = 3; i <=n; i++) {
        dp[i] = dp[i-1] + dp[i-2]; // Two base cases, so the total solution to step i will be the sum of two solutions (the step to i - 1 and step to i - 2)
    }
    return dp[n];
};
