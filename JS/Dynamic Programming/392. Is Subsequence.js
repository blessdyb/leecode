/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isSubsequence = function(s, t) {   // Compare two sequence, two pointer solution always be the first thing in your mind
    if (s.length > t.length) {
        return false; 
    }
    let i = 0;
    let j = 0;
    while(i < s.length && j < t.length) {
        if (s[i] === t[j]) {
            i++;
        }
        j++;
    }
    return s.length <= i;
};

var isSubsequence = function(s, t) {
/**
 *  DP: a. Base case: 
 *         i.  dp[0][j] = true, when s is empty, always return true
 *         ii. dp[i][0] = false, when t is empty, always return false
 *      b. State transform equation:
 *         i.  When s[i] == t[j],  dp[i][j] = dp[i - 1][j - 1]
 *         ii. When s[i] !== t[j],  dp[i][j] = dp[i][j - 1]
 **/
    const dp  = [];
    for(let i = 0; i <= s.length; i++) {
        dp[i] = [];
        for(let j = 0; j <= t.length; j++) {
            if (i === 0) {
                dp[i][j] = true;    
            } else if (j === 0) {
                dp[i][j] = false;
            } else if (s[i] === t[j]) {
                dp[i][j] = dp[i - 1][j - 1];
            } else {
                dp[i][j] = dp[i][j - 1];   
            }
        }
    }
    return dp[s.length][t.length];
}
