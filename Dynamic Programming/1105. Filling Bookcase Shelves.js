/**
 * @param {number[][]} books
 * @param {number} shelf_width
 * @return {number}
 */
var minHeightShelves = function(books, shelf_width) {
    const n = books.length;
    const cache = [];
    function helper(i) {
        if(i < n) {
            if (cache[i] === undefined) {
                let minHeight = Infinity;
                let leftWidth = shelf_width;
                let maxHeightInThisLayer = 0;
                let j = i;
                while(leftWidth > 0 && j < books.length && (leftWidth - books[j][0]) >= 0) {
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
