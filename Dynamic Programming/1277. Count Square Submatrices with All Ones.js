/**
 * @param {number[][]} matrix
 * @return {number}
 */
var countSquares = function(matrix) {
    const [m, n] = [matrix.length, matrix[0].length];
    let result = 0;
    const dp = []; // dp[i][j] means the number of squres with matrix[i][j] as the bottom-right vertex
    for(let i = 0; i < m; i++) {
        dp[i][j] = [];
        for(let j = 0; j < n; j++) {
           if (i === 0 || j === 0) {
              dp[i][j] = matrix[i][j]; 
           } else if (matrix[i][j] === 0) {
              dp[i][j] = 0;
           }  else {
              dp[i][j] = Math.min(dp[i - 1][j - 1], dp[i - 1][j], dp[i][j - 1]) + 1; 
           }
           result += dp[i][j];
        }
    }
    return result;
};
