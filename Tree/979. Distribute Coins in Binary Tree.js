/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
// How many coins need to pass through one node depends on given node's child node. So we can consider DFS to get the leaf node requirements, then populate to upper layer.
var distributeCoins = function(root) {
    let result = 0;
    (function dfs(node) {
        if (node) {
            const left = dfs(node.left);
            const right = dfs(node.right);
            result += Math.abs(left) + Math.abs(right);
            return node.val + left + right - 1; // Each node will need or give out Math.abs(node.val - 1) coins. 
        }
        return 0;
    })(root);
    return result;
};
