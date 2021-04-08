/**
 * @param {number} num
 * @return {number[]}
 */
var countBits = function(num) {
    const dp = [0];
    for(let i = 1; i <= num; i++) {
        if (i % 2 === 1) {
            dp[i] = dp[i - 1] + 1;
        } else {
            dp[i] = dp[Math.floor(i / 2)];
        }
    }
    return dp;
};

var countBits = function(num) {
    return (num & (n - 1)) + 1;   
}
