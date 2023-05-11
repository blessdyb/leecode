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
 * @param {number} low
 * @param {number} high
 * @return {number}
 */
var rangeSumBST = function(root, low, high) {
    let sum = 0;
    const stack = [root];
    while (stack.length) {
        const node = stack.pop();
        if (node) {
            if (node.val >= low && node.val <= high) {
                sum += node.val;
            }
            if (node.val > high) {
                stack.push(node.left);
            } else if (node.val < low) {
                stack.push(node.right);
            } else {
                stack.push(node.left);
                stack.push(node.right)
            }
        }
    }
    return sum;
};
