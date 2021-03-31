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
var deepestLeavesSum = function(root) {
    let result = 0;
    const stack = [root];
    let count = 1;
    let current = 0;
    while(stack.length) {
        const node = stack.shift();
        current++;
        if (node) {
            result += node.val;
            if (node.left) {
                stack.push(node.left);
            }
            if (node.right) {
                stack.push(node.right);
            }
            if (current === count) {
                count = stack.length;
                current = 0;
                if (count !== 0) {
                    result = 0;
                }
            }
        }
    }
    return result;
};
