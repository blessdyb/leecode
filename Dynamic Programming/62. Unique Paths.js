/**
 * @param {number} m
 * @param {number} n
 * @return {number}
 */
var uniquePaths = function(m, n) {
    const dp = Array(m + 1).fill(0).map(i => Array(n + 1).fill(0));
    dp[0][0] = 1;
    dp[0][1] = 1;
    dp[1][0] = 1;
    for(let i = 1; i <= m; i++) {
        for(let j = 1; j <=n; j++) {
            if (i === 1) {    // Note the special case, since robot can only go right or down, so for [m, 1] or [1, n] the result will always be 1
                dp[i][j] = 1;
            } else if (j === 1) {
                dp[i][j] = 1;
            } else {
                dp[i][j] = dp[i][j - 1] + dp[i - 1][j];   
            }
        }
    }
    return dp[m][n]
};
