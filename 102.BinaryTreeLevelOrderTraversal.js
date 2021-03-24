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
 * @return {number[][]}
 */
var levelOrder = function(root) {
    const stack = [root];
    let count = 1;
    let current = 0;
    let depth = 0;
    const result = [];
    while(stack.length) {
        const node = stack.shift();
        if (node) {
            if (current === 0) {
                result.push([]);
            } 
            result[depth].push(node.val);
            if (node.left) {
                stack.push(node.left);
            }
            if (node.right) {
                stack.push(node.right)
            }
            current++;
            if (current === count) {
                count = stack.length;
                current = 0;
                depth++;
            }
            
        }
    }
    return result;
};
