/**
 * @param {string} text1
 * @param {string} text2
 * @return {number}
 */
var longestCommonSubsequence = function(text1, text2) {
    const m = text1.length;
    const n = text2.length;
    const dp = new Array(m + 1).fill(0).map(i => new Array(n + 1).fill(0));
    for(let i = 1; i <= m; i++) {
        for(let j = 1; j <= n; j++) {
            if (text1[i - 1] === text2[j - 1]) {
                dp[i][j] = 1 + dp[i - 1][j - 1]; 
            } else {
                dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1]);
            }
        }
    }
    return dp[i][j];
}

var longestCommonSubsequence = function(text1, text2) { // Based on the state transform equation, we can reduce the memory usage by using a rolling array
  const m = text1.length;
  const n = text2.length;
  let dp = new Array(n + 1).fill(0);
  for(let i = 1; i <= m; i++) {
      const temp = new Array(n + 1).fill(0);
      for(let j = 1; j <= n; j++) {
          if (text1[i] === text2[j]) {
              temp[j] = 1 + dp[j - 1];
          } else {
              temp[j] = Math.max(temp[j - 1], dp[j]); 
          }
      }
      dp = temp;
  }
  return dp[n];
}
