/**
 * @param {number} n
 * @return {number}
 */
var countVowelStrings = function(n) {  // Recursive
    return (function counter(previousIndex, leftDigits) {
        if (leftDigits > 0) {
            let count = 0;
            for(let i = 0; i < 5; i++) {
                 count += counter(i, leftDigits - 1);
            }
            return count;
        }
        return 1; //  base case when n === 0, it means we have assigned the string, so the value is 1.
    })(0, n);
};

var countVowelStrings = function(n) {  // DP  dp[i][j] = Sum(dp[i - 1][?]) = dp[i - 1][j] + dp[i][j + 1]
    const dp = [[1, 1, 1, 1, 1]];
    for(let i = 1; i <=n; i++) {
        dp[i] = [];
        dp[i][4] = 1;  // given n length you can only have one combination of string with all characters as u
        for (let j = 3; j >= 0; j--) {
            dp[i][j] = dp[i - 1][j] + dp[i][j + 1];   
        }
    }
    return dp[n][0];
}

var countVowelStrings = function(n) {
    let [a, e, i, o, u] = [1, 1, 1, 1, 1];
    while(n - 1 > 0) {
        u = u;
        o += u;
        i += o;
        e += i;
        a += e;
        n--;
    }
    return a + e + i + o + u;
};

