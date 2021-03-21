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
var minDiffInBST = function(root) {
    let min = Infinity;
    const stack = [];
    let current = root;
    let previous;
    while(true) {
        if (current) {
            stack.push(current);
            current = current.left;
        } else if (stack.length) {
            const node = stack.pop();
            if (node && previous) {
                min = Math.min(min, Math.abs(previous.val - node.val)); 
            }
            previous = node;
            current = node.right;
        } else {
            break;
    }
    return min;
};
