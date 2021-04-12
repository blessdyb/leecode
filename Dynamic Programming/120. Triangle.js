/**
 * @param {number[][]} triangle
 * @return {number}
 */
var minimumTotal = function(triangle) {  // DP
    const dp = [[triangle[0][0]]];
    const n = triangle.length;
    for(let i = 1; i < n; i++) {
        dp[i] = [triangle[i][0] + dp[i - 1][0]];
        dp[i][i] = triangle[i][i] + dp[i - 1][i - 1];
        for(let j = 1; j < i; j++) {
            dp[i][j] = triangle[i][j] + Math.min(dp[i - 1][j], dp[i - 1][j - 1]);  
        }
    }
    return Math.min(...dp[n - 1]);
};
