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
    return (function(start, end) {
        const result = [];
        for(let i = start; i <= end; i++) {
            generateTrees(start, i - 1).forEach(left => {
                generateTrees(i + 1, end).forEach(right => {
                    const node = new TreeNode(i, left, right);
                    result.push(node);
                });
            });
        }
        return result.length ? result : [null];
    })(1, n);
};

