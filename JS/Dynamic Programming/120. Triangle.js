/**
 * @param {number[][]} triangle
 * @return {number}
 */
var minimumTotal = function(triangle) {  // DP from top to bottom
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

var minimumTotal = function(triangle) {  // DP from bottom to top
    const n = triangle.length;
    const dp = triangle[n - 1];
    for(let i = n - 2; i >= 0; i--) {
        for(let j = 0; j <= i; j++) {
            dp[j] = triangle[i][j] + Math.min(dp[j], dp[j + 1]);   
        }
    }
    return dp[0];
}
