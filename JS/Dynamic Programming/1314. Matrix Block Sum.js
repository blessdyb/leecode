/**
 * @param {number[][]} mat
 * @param {number} k
 * @return {number[][]}
 */
var matrixBlockSum = function(mat, k) {  // 2D Array query / preSum
    const dp = [];
    const [m, n] = [mat.length, mat[0].length];
    for(let i = 0; i <= m; i++) {
        dp[i] = [];
        for(let j = 0; j <= n; j++) {
            if (i === 0 || j === 0) {
                dp[i][j] = 0; 
            } else {
                dp[i][j] = mat[i - 1][j - 1] + dp[i - 1][j] + dp[i][j - 1] - dp[i - 1][j - 1]; 
            }
        }
    }
    const answer = [];
    for(let i = 0; i < m; i++) {
        answer[i] = [];
        for(let j = 0; j < n; j++) {
            const [x1, y1] = [Math.max(0, i - k), Math.max(0, j - k)];
            const [x2, y2] = [Math.min(m - 1, i + k), Math.min(n - 1, j + k)];
            answer[i][j] = dp[x2 + 1][y2 + 1] + dp[x1][y1] - dp[x1][y2 + 1] - dp[x2 + 1][y1]; 
        }
    }
    return answer;
};
