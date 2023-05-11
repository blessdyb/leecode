/**
 * @param {number[]} cost
 * @return {number}
 */
var minCostClimbingStairs = function(cost) { // DP
    const dp = [0, 0]; // Since you can only take one / two steps, so dp[1] = 0, dp[2] = Math.min(cost[0], cost[1]));
    for (let i = 2; i <= cost.length; i++) {
        dp[i] = Math.min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2]);
    }
    return dp[cost.length];
};

var minCostClimbingStairs = function(cost) {
    let i = 0;
    let j = 0;
    for(let k = 2; k <= cost.length; k++) { // Note: we need to pass the last element of cost, so here we use <= instead of <
        const temp = Math.min(i + cost[k - 2], j + cost[k - 1]);
        [i, j] = [j, temp];
    }
    return j;
}
