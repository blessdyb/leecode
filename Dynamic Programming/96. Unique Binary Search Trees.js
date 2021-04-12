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
