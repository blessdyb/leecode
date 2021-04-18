/**
 * @param {number[]} piles
 * @return {boolean}
 */
var stoneGame = function(piles) {
    // Since for the base case when there are two piles, who plays first will always win.
    return true;
};

var stoneGame = function(piles) {  // Cache + Recursive
    const cache = {};
    return (function helper(i, j) {
        if (i > j) {
            return 0; 
        } else {
            const key = `${i}_${j}`;
            if (cache[key] === undefined) {
                if ((j - i) % 2 === 1) {   // First player
                    cache[key] = Math.max(pile[i] + helper(i + 1, j), pile[j] + helper(i, j - 1));
                } else {
                    cache[key] = Math.min(-pile[i] + helper(i + 1, j), -pile[j] + helper(i, j - 1)); 
                }
            }
            return cache[key];
        }
    })(0, piles.length - 1) > 0;
};
