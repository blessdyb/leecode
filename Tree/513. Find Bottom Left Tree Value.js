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
var findBottomLeftValue = function(root) {    // Bottom left node means we need to traverse the given tree with BSF way and get the first one in the bottom layer
    const stack = [root];
    let value;
    let current = 0;
    let count = 1;
    while(stack.length) {
        const node = stack.shift();
        if (node) {
            if (current === 0) {
                value = node.val;
            }
            current++;
            if (node.left) {
                stack.push(node.left);
            }
            if (node.right) {
                stack.push(node.right);
            }
            if (current === count) {
                current = 0;
                count = stack.length;
            }
        }
    }
    return value;
};
