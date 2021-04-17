/**
 * @param {number[]} arr
 * @param {number} k
 * @return {number}
 */
var maxSumAfterPartitioning = function(arr, k) {  // DP, dp[i] = dp[i - j] + j * Max(i - 1 ... i - j)
    const n = arr.length;
    const dp = [0];
    for(let i = 1; i <= n; i++) {
        let max = -Infinity;
        let best = 0;
        for(let j = 1; j <= k && i - j >= 0; j++) {
            max = Math.max(max, arr[i - j]);
            best = Math.max(best, dp[i - j] + max * j);
        }
        dp[i] = best;
    }
    return dp[n];
};
