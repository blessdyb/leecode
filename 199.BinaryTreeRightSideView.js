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
var rightSideView = function(root) {
    const stack = [root];
    let count = 1;
    let current = 0;
    const result = [];
    while(stack.length) {
        const node = stack.shift();
        current++;
        if (node) {
            if (node.left) {
                stack.push(node.left);
            }
            if (node.right) {
                stack.push(node.right);
            }
            if (count === current) {
                result.push(node.val);
                current = 0;
                count = stack.length;
            }
        }
    }
    return result;
};
