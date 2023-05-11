/**
 * @param {number[][]} books
 * @param {number} shelf_width
 * @return {number}
 */
var minHeightShelves = function(books, shelf_width) { // Cache + recursive
    const n = books.length;
    const cache = [];
    function helper(i) {
        if(i < n) {
            if (cache[i] === undefined) {
                let minHeight = Infinity;
                let leftWidth = shelf_width;
                let maxHeightInThisLayer = 0;
                let j = i;
                while(leftWidth > 0 && j < books.length && (leftWidth - books[j][0]) >= 0) { // Try all possible combinations
                    maxHeightInThisLayer = Math.max(maxHeightInThisLayer, books[j][1])
                    minHeight = Math.min(minHeight, maxHeightInThisLayer + helper(j + 1));
                    leftWidth -= books[j][0];
                    j++;
                }
                cache[i] = minHeight;
            }
            return cache[i];
        }
        return 0;
    }
    return helper(0);
};

/**
 * @param {number[][]} books
 * @param {number} shelf_width
 * @return {number}
 */
var minHeightShelves = function(books, shelf_width) {
    const n = books.length;
    const dp = [0];
    for(let i = 1; i <= n; i++) {
        let leftWidth = shelf_width;
        let maxHeightInLayer = 0;
        for(let j = i; j <= n; j++) {
            if (leftWidth >= 0 && leftWidth - books[j - 1][0] >= 0) {
                leftWidth -= books[j - 1][0];
                maxHeightInLayer = Math.max(maxHeightInLayer, books[j - 1][1]);
                dp[j] = Math.min(dp[j] || Infinity, dp[i - 1] + maxHeightInLayer);  // From layer one, put try to all books in one layer vs all other combinations
            } else {
                break;   
            }
        }
    }
    return dp[n];
};
