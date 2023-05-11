/**
 * @param {number[]} days
 * @param {number[]} costs
 * @return {number}
 */
var mincostTickets = function(days, costs) {
    const n = days.length;
    const cache = {};
    function getIndex(start, ticket) {
        let result = n; // if the ticket can cover the rest days, then we don't need to buy anymore tickets
        const expiredDay = days[start] + ticket;
        for(let i = start + 1; i < n; i++) {
            if (days[i] >= expiredDay) {
                result = i;
                break;
            }
        }
        return result;
    }
    return (function helper(start) {
        if (start >= n) {
            return 0; 
        }
        if (cache[start] === undefined) {
            const oneDay = costs[0] + helper(start + 1);
            const sevenDays = costs[1] + helper(getIndex(start, 7));
            const thirtyDays = costs[2] + helper(getIndex(start, 30));        
            cache[start] = Math.min(oneDay, sevenDays, thirtyDays);
        }
        return cache[start];
    })(0);
}; 

var mincostTickets = function(days, costs) {
    const dp = [0];
    for(let i = 1; i <= 365; i++) {
        if (days.indexOf(i) === -1) {
            dp[i] = dp[i - 1];   
        } else {
            dp[i] = Math.min(costs[0] + dp[Math.max(0, i - 1)], costs[1] + dp[Math.max(0, i - 7)], costs[2] + dp[Math.max(0, i - 30)]);   
        }
    }
    return dp[365];
};
