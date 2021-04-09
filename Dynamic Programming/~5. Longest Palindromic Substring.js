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
