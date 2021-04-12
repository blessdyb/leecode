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

var numDecodings = function(s) {  // State transform equation dp[i] = dp[i + 1] + dp[i + 2] conditional
    const dp = [];
    const n = s.length;
    dp[n] = 1;
    if (s[n - 1] === '0') {
        dp[n - 1] = 0;    
    } else {
        dp[n - 1] = 1;   
    }
    for (let i = n - 2; i >= 0; i--) {
        if (s[i] === '0') {
            dp[i] = 0;   
        } else {
            const firstDigit = dp[i + 1];
            let secondDigits = 0;
            if (parseInt(s.substr(i, 2), 10) < 27) {
                secondDigits = dp[i + 2];   
            }
            dp[i] = firstDigit + secondDigits;
        }
    }
    return dp[0];
};

var numDecodings = function(s) {  // DP + Space compression
    let second = 1;
    let first = 0;
    const n = s.length;
    if (s[n - 1] !== '0') {
        first = 1;   
    }
    for (let i = n - 2; i >= 0; i--) {
        if (s[i] === '0') {
            second = first;
            first = 0;
        } else {
            let temp = parseInt(s.substr(i, 2), 10) < 27 ? second : 0;
            second = first;
            first = first + temp;
        }
    }
    return first;
}
