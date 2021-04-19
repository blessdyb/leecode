/**
 * @param {number} n
 * @return {number}
 */
var nthPersonGetsNthSeat = function(n) {
    return n === 1 ? 1 : 0.5;  
};

var nthPersonGetsNthSeat = function(n) {
    const dp = [0, 1, 0.5];
    for (let i = 3; i <= n; i++) {
        dp[i] = 1 / i + (i - 2) * 1 / i * dp[i - 1];
    }
    return dp[n];
};
