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
var sumNumbers = function(root) {
    const stack = [[root, '']];
    let depth = 0;
    let result = 0;
    while(stack.length) {
        const [node, parent] = stack.shift();
        if (node) {
            const value = parent + node.val;
            if (!node.left && !node.right) {
                result += parseInt(value, 10);
            } else {
                if (node.left) {
                    stack.push([node.left, value]);
                }
                if (node.right) {
                    stack.push([node.right, value]);
                }
            }
        }
    }
    return result;
};
