/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function(n) {
    const dp = [0, 1, 2];
    for(let i = 3; i <=n; i++) {
        dp[i] = dp[i-1] + dp[i-2]; // Two base cases, so the total solution to step i will be the sum of two solutions (the step to i - 1 and step to i - 2)
    }
    return dp[n];
};


var climbStairs = function(n) {
    if (n === 1) {
        return 1;   
    } else if (n === 2) {
        return 2;   
    } else {
        let first = 1;     // Since we only care about the result, we can save some space with Fibonacci
        let second = 2;
        for(let i = 3; i <=n; i++) {
            [first, second] = [second, first];
            second = first + second;
        }
        return second;
    }
}
