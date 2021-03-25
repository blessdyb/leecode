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
var increasingBST = function(root) {
    let result = null;
    const stack = [];
    let current = root;
    let pointer;
    while(true) {
        if (current) {
            stack.push(current);
            current = current.left;
        } else if (stack.length) {
            const node = stack.pop();
            if (!result) {
                result = pointer = new TreeNode(node.val);  
            } else {
                pointer.left = new TreeNode(node.val);  
                pointer = pointer.left;
            }
            current = node.right;
        } else {
            break;
        }
    }
    return result;
};
