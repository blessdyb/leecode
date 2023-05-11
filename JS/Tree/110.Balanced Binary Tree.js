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
 * @return {boolean}
 */
const maxLength = function(node) {
    if(node) {
        return 1 + Math.max(maxLength(node.left), maxLength(node.right));
    }
    return 0;
}
var isBalanced = function(root) {
    if (!root || (!root.left && !root.right)) {
        return true; 
    }
    return Math.abs(maxLength(root.left) - maxLength(root.right)) < 2 && isBalanced(root.left) && isBalanced(root.right);
};
