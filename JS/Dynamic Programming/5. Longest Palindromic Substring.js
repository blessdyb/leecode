/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function(s) {          // Brute force
    let result = '';
    let max = 0;
  
    function isPalindorm(s) {
        for(let i = 0; i < s.length; i++) {
            if (s[i] !== s[s.length - 1 - i]) {
                return false; 
            }
        }
        return true;
    }
    
    for(let i = 0; i < s.length; i++) {
        for(let j = i + 1; j < s.length; j++) {
            if (isPalindrom(s.substr(i, j)) && (j - i > max)) {
                max = j - i;
                result = s.substr(i, j);
            }
        }
    }
    return result;
};

var longestPalindrom = function(s) {  // DP,  dp[i][j] = dp[i + 1][j - 1] && s[i] === s[j]
    const dp = [];
    const n = s.length;
    let maxLen = 0;
    let result = '';
    for(let i = 1; i <= n; i++) {  // Since the state transfer equation tells us that long string's value relys on shorter one's. We iterate all different sizes from 1 ~ n
        for(let start = 0; start < i; start++) {
            const end = i - 1 - start;
            if (end >= n) {
                break;   
            }
            if (i <= 2) {
                dp[start][end] = s[start] === s[end];   
            } else {
                dp[start][end] = dp[start + 1][end - 1] && s[start] === s[end];
            }
            if (dp[start][end] && i > maxLen) {
                result = s.substr(start, i);
                maxLen = i;
            }
        }
    }
    return result;
};

var longestPalindrome = function(s) { // Base on the state transform equation, we can reduce the space usage
    const dp  = [];
    const n = s.length;
    let max = 0;
    let result = '';
    for(let i = n - 1; i >= 0; i--) {
        for(let j = n - 1; j >= i; j--) {
            if (j - i < 3) {
                dp[j] = s[i] === s[j];   
            } else {
                dp[j] = s[i] === s[j] && dp[j - 1];
            }
            if (dp[j] && j - i + 1 > max) {
                result = s.substr(i, j - i + 1);
                max = j - i + 1;
            }
        }
    }
    return max;
};

var longestPalindrome = function(s) {
    let left = 0, right = 0;
    const n = s.length;
    
    function expandFromCenter(l, r) {
        while(l >= 0 && r < n && s[l] === s[r]) {
            l--;
            r++;
        }
        return r - l - 1;
    }
    
    for(let i = 0; i < n; i++) {
        let len = Math.max(expandFromCenter(i, i), expandFromCenter(i, i + 1));
        if (len > right - left) {
            left = i + Math.floor((len - 1) / 2));
            right = i + Math.floor(len / 2);
        }
    }
    return s.substr(left, right - left + 1);
}
