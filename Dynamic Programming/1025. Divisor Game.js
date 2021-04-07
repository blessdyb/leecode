/**
 * @param {number} n
 * @return {boolean}
 */
var divisorGame = function(n) {                   // DP
    const dp = Array(n + 1).fill(false);          // Base case: dp[0] = false & dp[1] = true
    dp[1] = true;
    for (let i = 2; i <= n; i++) {
        for (let j = 1; j < i; j++) {
            if (i % j === 0 && !dp[i - j]) {      // State transform equation, can make a move 0 < j < i && i % j === 0 and sub problem dp[i - j] is false
                dp[i] = true;
                break;
            }
        }
    }
    return dp[n];
};
