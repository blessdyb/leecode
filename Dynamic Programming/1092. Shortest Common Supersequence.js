/**
 * @param {string} str1
 * @param {string} str2
 * @return {string}
 */
var shortestCommonSupersequence = function(str1, str2) {
  // Get LCS
  const m = str1.length;
  const n = str2.length;
  const dp =  new Array(m + 1).fill(0).map(item => new Array(n + 1).fill(0));
  for(let i = 1; i <= m; i++) {
      for(let j = 1; j <= n; j++) {
          if (str1[i - 1] === str2[j - 1]) {
              dp[i][j] = 1 + dp[i - 1][j - 1]; 
          } else {
              dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1]); 
          }
      }
  }
  let i = m, j = n, result = '';
  while(i || j) {
      if (i === 0) {
          result = str2.slice(0, j) + result; 
          break;
      } else if (j === 0) {
          result = str1.slice(0, i) + result;
          break;
      } else if (dp[i][j] > dp[i - 1][j] && dp[i][j] > dp[i][j - 1]) {
          i--;
          j--;
          result = str1[i] + result;
      } else if (dp[i][j] === dp[i - 1][j]) {
          i--;
          result = str1[i] + result;
      } else {
          j--;
          result = str2[j] + result; 
      }
  }
  return result;
}
