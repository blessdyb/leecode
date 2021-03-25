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
var averageOfLevels = function(root) {
    let result = [];
    let stack = [root];
    let count = 1;
    let current = 0;
    let temp = 0;
    while(stack.length) {
        const n = stack.shift();
        if (n) {
            current++;
            temp += n.val;
            if (n.left) {
                stack.push(n.left);
            }
            if (n.right) {
                stack.push(n.right);
            }
            if (current === count) {
                result.push(temp / count);
                temp = 0;
                count = stack.length;
                current = 0;
            }
        }
    }
    return result;
};
