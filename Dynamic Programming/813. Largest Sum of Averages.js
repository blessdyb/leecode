/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var largestSumOfAverages = function(nums, k) {
    const n = nums.length;
    const dp = new Array(k + 1).fill(0).map(item => new Array(n + 1).fill(0));
    const sums = new Array(n + 1).fill(0);
    for(let i = 1; i <= n; i++) {
        sums[i] = sums[i - 1] + nums[i - 1];
    }
  
    // Recursive solution,  dp[n][k] = Max(dp[0..i][k - 1] + Avg(nums[i..n]))
    // For given problem [nums, k],  we can get the sub-problem [nums[0..i], k - 1] + Avg(nums[i..n]). 
    // It means we find a position to split nums to two group, the last half one we treat it as one set and get the average
    // The first half which can be split into at most k - 1 sets (the sub-problem)
    return (function lsa(n, k) {
        if (dp[k][n] === 0) {
            if (k === 1) {
                dp[k][n] = sums[n] / n; 
            } else {
                for(let i = k - 1; i < n; i++) {
                    dp[k][n] = Math.max(dp[k][n], lsa(i, k - 1) + (sums[n] - sums[i]) / (n - i)); 
                }
            }
        }
        return dp[k][n];
    })(n, k);
}
