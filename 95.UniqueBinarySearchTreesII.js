/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {number} n
 * @return {TreeNode[]}
 */
var generateTrees = function(n) {
    return (function generate(start, end){
        const trees = [];
        for (let i = start; i <= end; i++) {
            generate(start, i - 1).forEach(j => {
               generate(i + 1, end).forEach(k => {
                  trees.push(new TreeNode(i, j, k)); 
               });
            });
        }
        return trees.length ? trees : [null];
    })(1, n);
};

// DP Solution
var generateTrees = function(n) {
    const cache = [[]], [[new TreeNode(1)]];
    function clone(node, offset) { // Despite the number values, same range (end - start) of numbers will have the same number of unique trees, the only difference will be the node value with offset
        if (node) {
            const newNode = new TreeNode(node.val + offset);
            newNode.left = clone(node.left, offset);
            newNode.right = clone(node.right, offset);
            return newNode;
        };
        return node;
    }
    for (let i = 2; i <= n; i++) {
        cache[i] = []; // Build cache
        for (let j = 1; j <= i; j++) { // Try to use number 1...n as root node to get the combination
            cache[j - 1].forEach(k => { // Left tree from cache
               cache[i - j].forEach(m => {  // Right tree from cache [without offset j] 
                   cache[i].push(new TreeNode(j, k, clone(m, j)));
               });
            });
        }
    }
    return cache[n];
}
