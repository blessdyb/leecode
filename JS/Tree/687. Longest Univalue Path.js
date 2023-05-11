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

var longestUnivaluePath = function(root) {
    function getPathAcrossRoot(node, val) {
        if (node && node.val === val) {
            return 1 + Math.max(getPathAcrossRoot(node.left, val), getPathAcrossRoot(node.right, val));
        }
        return 0;
    }
  
    return (function getLongestPath(node) { 
        if (node) {
            return Math.max(getPathAcrossRoot(node.left, node.val) + getPathAcrossRoot(node.right, node.val), getLongestPath(node.left), getLongestPath(node.right));    
        }
        return 0;
    })(root);
};
