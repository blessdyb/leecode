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
    const stack = [root];
    let count = 1;
    let current = 0;
    let previous;
    while(stack.length) {
        const node = stack.shift();
        current++;
        if (node) {
            if (previous) {
                previous.next = node;
            }
            previous = node;
            if (node.left) {
                stack.push(node.left);
            }
            if (node.right) {
                stack.push(node.right);
            }
            if (count === current) {
                count = stack.length;
                current = 0;
                previous = null;
            }
        }
    }
    return root;
};

