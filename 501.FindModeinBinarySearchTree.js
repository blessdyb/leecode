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
var findMode = function(root) {
    let result = [];
    let current = 0;
    let count = 0;
    let maxCount = 0;
    (function dfs(node) {
        if (node) {
            dfs(node.left);
            if (node.val != current) {
                current = node.val;
                count = 0;
            }
            count++;
            if (maxCount === count) {
                result.push(current);  
            } else if (maxCount < count) {
                result = [current];
                maxCount = current;
            }
            dfs(node.right);
        }
    })(root);
    return result;
};
