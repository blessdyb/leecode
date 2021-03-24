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
var isEvenOddTree = function(root) {
    const stack = [root];
    let count = 1;
    let current = 0;
    let depth = 0;
    let previous = null;
    while(stack.length) {
        const node = stack.shift();
        if (depth % 2 === 0) {
            if (node.val % 2 === 0) {
                return false;
            }
            if (previous !== null && previous >= node.val) {
                return false;
            }
        } else {
            if (node.val %2 === 1) {
                return false;
            }
            if (previous !== null && previous <= node.val) {
                return false;
            }
        }
        current++;
        previous = node.val;
        if (node.left) {
            stack.push(node.left);
        }
        if (node.right) {
            stack.push(node.right);
        }
        if (current === count) {
            current = 0;
            depth++;
            count = stack.length;
            previous = null;
        }
    }
    return true;
};
