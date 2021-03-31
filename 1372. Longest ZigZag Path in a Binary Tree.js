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

var longestZigZag = function(root) {
    let max = -Infinity;
    (function zigzag(node) {
        if (node) {
            let [leftLength1, rightLength1] = zigzag(node.left);
            let [leftLength2, rightLength2] = zigzag(node.right);
            max = Math.max(max, rightLength1 + 1, leftLength2 + 1); // Bottom up from leaf nodes to root. We compare each node's left subtree's right subtree to its right subtree's left subtree.
            return [rightLength1 + 1, leftLength2 + 1];
        }
        return [0, 0];
    })(root);
    return max - 1;
};
