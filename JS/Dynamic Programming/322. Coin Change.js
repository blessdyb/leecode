/**
 * @param {number[]} coins
 * @param {number} amount
 * @return {number}
 */
var coinChange = function(coins, amount) {
    const dp = [0];
    const minCoin = Math.min(...coins);
    for(let i = 1; i <= amount; i++) {
        if (i < minCoin) {
            dp[i] = -1; 
        } else if (coins.indexOf(i) !== -1) {
            dp[i] = 1; 
        } else {
            let hasSolution = false;
            let min = Infinity;
            for(let j = 0; j < coins.length; j++) {
                const temp = i - coins[j];
                if (temp >= minCoin && dp[temp] !== -1) {
                    hasSolution = true;
                    min = Math.min(min, dp[temp] + 1)
                }
            }
            dp[i] = hasSolution ? min : -1;
        }
    }
    return dp[amount];
};
