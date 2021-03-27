/**
 * // Definition for a Node.
 * function Node(val, left, right, next) {
 *    this.val = val === undefined ? null : val;
 *    this.left = left === undefined ? null : left;
 *    this.right = right === undefined ? null : right;
 *    this.next = next === undefined ? null : next;
 * };
 */

/**
 * @param {Node} root
 * @return {Node}
 */
var connect = function(root) {
    if (root) {
        (function c(left, right) {
            if (left && right) {
                left.next = right;
                c(left.left, left.right)
                c(left.right, right.left)
                c(right.left, right.right)
            }
        })(root.left, root.right);
    }
    return root;
};
