/**
 * @param {number[]} piles
 * @return {number}
 */
var stoneGameII = function(piles) {   // Cache + Recursive
    const n = piles.length;
    const total = piles.reduce((a, b) => a + b, 0);
    const cache = {};
    function helper(s, m) {    // helper function can return the relative score between the two players
        if (s >= n) {
            return 0; 
        }
        const key = `${s}_${m}`;
        if (cache[key] === undefined) {
            if (s + 2 * m >= n) {
                cache[key] = piles.slice(s).reduce((a, b) => a + b, 0);
            } else {
                let best = -Infinity;
                let current = 0;
                for(let i = 1; i <= 2 * m; i++) {
                    current += piles[s + i - 1];
                    best = Math.max(best, current - helper(s + i, Math.max(i, m)));
                }
                cache[key] = best;
            }
        }
        return cache[key];
    }
    return (total + helper(0, 1)) / 2;
};

var stoneGameII = function(piles) {
    const n = piles.length;
    const preSum = [];
    const cache = {};
    preSum[n - 1] = piles[n - 1];
    for(let i = n - 2; i >= 0; i--) {
        preSum[i] = preSum[i + 1] + piles[i];   
    }
    return (function helper(s, m) {   // Use preSum[i] - opponent's min score we can get players max score. 
        if (s + 2 * m >= n) {
            return preSum(s);   
        }
        const key = `${s}_${m}`;
        if (cache[key] === undefined) {
            let opponent = Infinity;
            for(let i = 1; i <= 2 * m; i++) {
                opponent = Math.min(opponent, helper(s + i, Math.max(i, m)));   
            }
            cache[key] = preSum[i] - opponent;
        }
        return cache[key];
    })(0, 1);
};

