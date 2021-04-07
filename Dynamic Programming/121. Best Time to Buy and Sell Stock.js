/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {    // Max profile only happens when we buy at a lower price and then sell at a higher price. Means we need to find a peak after the lowest point.
    let lowestPrice = Infinity;
    let profit = 0;
    for (let i = 0; i < prices.length; i ++) {
        if (lowestPrice > prices[i]) {
            lowestPrice = prices[i];    // Get the lowest price
        } else {
            profit = Math.max(profit, prices[i] - lowestPrice);  // Check all the diff between lowest price and the prices beyond. 
        }
    }
    return profit;
};


var maxProfit = function(prices) {
    const dp = [[0, -prices[0]]];    // Base case, 0 day [no profile if we do nothing, negative price if we just buy].
    for (let i = 1; i < prices.length; i++) {
        dp[i] = [];
        dp[i][0] = Math.max(dp[i - 1][0], dp[i - 1][1] + prices[i]);
        dp[i][1] = Math.max(dp[i - 1][1], -prices[i]);
    }
    return dp[prices.length - 1][0];   // Result we just need to check when we sell stock, if we keep stock in hand, the profit is negative.
}
