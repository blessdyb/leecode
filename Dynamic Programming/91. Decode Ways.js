/**
 * @param {string} s
 * @return {number}
 */
var numDecodings = function(s) {  // Recursive
    return (function decode(n) {
        if (n === s.length) {
            return 1; 
        }
        if (s[n] === '0') {
            return 0; 
        }
        const singleDigit = decode(n + 1);
        let doubleDigits = 0;
        if (n < s.length - 1) {
            if (parseInt(s.substr(n, 2), 10) < 27) {
                doubleDigits = decode(n + 2); 
            }
        }
        return singleDigit + doubleDigits;
    })(1);
};

var numDecodings = function(s) {  // Recursive + cache
    const dp = [];
    return (function decode(n) {
        if (dp[n] !== undefined) {
            return dp[n]; 
        }
        if (n === s.length) {
            return 1; 
        }
        if (s[n] === '0') {
            return 0; 
        }
        const singleDigit = decode(n + 1);
        let doubleDigits = 0;
        if (n < s.length - 1) {
            if (parseInt(s.substr(n, 2), 10) < 27) {
                doubleDigits = decode(n + 2); 
            }
        }
        dp[n] = singleDigit + doubleDigits;
        return singleDigit + doubleDigits;
    })(1);
};
