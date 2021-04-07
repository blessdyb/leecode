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
 * @return {number[]}
 */
var largestValues = function(root) {
    const stack = [root];
    const result = [];
    let count = 1;
    let current = 0;
    let max = -Infinity;
    while(stack.length) {
        const node = stack.shift();
        if (node) {
            max = Math.max(max, node.val);
            current++;
            if (node.left) {
                stack.push(node.left);
            }
            if (node.right) {
                stack.push(node.right);
            }
            if (current === count) {
                result.push(max);
                max = -Infinity;
                count = stack.length;
                current = 0;
            }
        }
    }
    return result;
};
