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
 * @return {boolean}
 */
var isUnivalTree = function(root) {
    const stack = [root];
    let current = -1; // Make sure the initial value of current is not in the scope of the node value
    while(stack.length) {
        const node = stack.shift();
        if (node) {
            if (current === -1) {
                current = node.val; 
            }
            if (current !== node.val) {
                return false; 
            }
            stack.push(node.left);
            stack.push(node.right);
        }
    }
    return true;
};
