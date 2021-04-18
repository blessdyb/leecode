/**
 * @param {number[][]} matrix
 * @return {number}
 */
var minFallingPathSum = function(matrix) {  // Straightforward DP
    const [m, n] = [matrix.length, matrix[0].length];
    const dp = [matrix[0]];
    for(let i = 1; i < n; i++) {
        dp[i] = [];
        for(let j = 0; j < m; j++) {
            const left = Math.max(0, j - 1);
            const right = Math.min(m - 1, j + 1);
            dp[i][j] = matrix[i][j] + Math.min(dp[i - 1][j], dp[i - 1][left], dp[i - 1][right]);
        }
    }
    return Math.min(...dp[n - 1]);
};
