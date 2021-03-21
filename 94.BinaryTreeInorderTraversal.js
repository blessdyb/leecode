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
var inorderTraversal = function(root) {
    const stack = [];
    let current = root;
    const result = [];
    while(true) {
        if (current) {
            stack.push(current);
            current = current.left;
        } else if (stack.length) {
            const node = stack.pop();
            result.push(node.val);
            current = node.right;
        } else {
            break;
        }
    }
    return result
};
