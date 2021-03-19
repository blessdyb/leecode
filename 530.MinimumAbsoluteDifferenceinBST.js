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
// BST means it's a sorted array when tranversing with DFS, min absolute diff only happens in adjacent elements, so we just need to remember previous node when tranversing
var getMinimumDifference = function(root) {
    let min = Infinity;
    let prev;
    (function dfs(node){
        if (node) {
            dfs(node.left);
            if (prev) {
                min = Math.min(min, Math.abs(prev.val - node.val));   
            }
            prev = node;
            dfs(node.right);
        }
        return min;
    })(root);
    return min;
};
