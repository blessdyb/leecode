/**
 * @param {number} n
 * @return {number}
 */
var numTrees = function(n) {
    if (n < 2) {
        return 1; 
    } 
    let result = 0;
    for(let i = 0; i < n; i++) {
        result += numTrees(i) * numTrees(n - i - 1);
    }
    return result;
};

var numTrees = function(n) {
    const dp = [1, 1];
    for (let i = 2; i <=n; i++) {
        let temp = 0;
        for (let j = 0; j < i; j++) {
            temp += dp[j] * dp[i - j - 1];
        }
        dp[i] = temp;
    }
    return dp[n];
}
