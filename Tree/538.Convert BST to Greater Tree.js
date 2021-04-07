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
 * @return {TreeNode}
 */
var convertBST = function(root) {
    let sum = 0;
    const stack = [];
    let current = root;
    while(true) {
        if (current) {
            stack.push(current);
            current = current.right;
        } else if (stack.length) {
            const node = stack.pop();
            sum += node.val;
            node.val = sum;
            current = node.left;
        } else {
            break;
        }
    }
    return root;
};
