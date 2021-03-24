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
 * @return {number[][]}
 */
var zigzagLevelOrder = function(root) {
    const stack = [root];
    const result = [];
    let count = 1;
    let current = 0;
    let depth = 0;
    while(stack.length) {
        const node = stack.shift();
        if (node) {
            if (current === 0) {
                result.push([]);
            }
            result[depth][depth % 2 === 0 ? 'push' : 'unshift'](node.val);
            if (node.left) {
                stack.push(node.left);
            }
            if (node.right) {
                stack.push(node.right);
            }
            current++;
            if (current === count) {
                current = 0;
                count = stack.length;
                depth++;
            }
        }
    }
    return result;
};
