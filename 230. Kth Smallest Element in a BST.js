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
 * @param {number} k
 * @return {number}
 */
var kthSmallest = function(root, k) {           // DFS of a BST will get a sorted array, so the nth time we visit the node, the node will be the nth smallest
    let result;
    (function dfs(node) {
        if (node.left) {
            dfs(node.left);
        }
        k--;
        if (k === 0) {
            result = node.val;
        }
        if (node.right) {
            dfs(node.right);
        }
    })(root);
    return result;
};
