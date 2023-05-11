/**
 * @param {string} s
 * @return {number}
 */
var countSubstrings = function(s) {
    let result = 0;
    const n = s.length;
    const dp = [];
    for(let i = n - 1; i >= 0; i--) {
        dp[i] = [];
        for(let j = i; j < n; j++) {
            if(i === j) {
                dp[i][j] = true;
                result++;
            } else if (j - i === 1 && s[i] === s[j]) {
                dp[i][j] = true; 
                result++;
            }  else if (s[i] === s[j] && dp[i + 1][j - 1]) {
                dp[i][j] = true;
                result++;
            }
        }
    }
    return result;
};

var countSubstrings = function(s) {
    let result = 0;
    const n = s.length;
    
    function countPalindrom(i, j) {
        let count = 0;
        while(i >= 0 && j < n) {
            if (s[i] !== s[j]) {
                break; 
            }
            i--;
            j++;
            count++;
        }
        return count;
    }
  
    for(let i = 0; i < n; i++) {
        result += countPalindrom(i, i);
        result += countPalindrom(i, i + 1);
    }
    return result;
}
