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
var rob = function(root) {  // Recursively check the value of rob root node or not rob root node.
    if (root) {
        const robRoot = root.val + (root.left ? rob(root.left.left) + rob(root.left.right) : 0) + (root.right ? rob(root.right.left) + rob(root.right.right) : 0);
        const notRobRoot = rob(root.left) + rob(root.right);
        return Math.max(robRoot, notRobRoot);
    }
    return 0;   
};
