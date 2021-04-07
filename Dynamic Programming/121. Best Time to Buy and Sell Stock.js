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
