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
 * @param {number} targetSum
 * @return {number[][]}
 */
var pathSum = function(root, targetSum) {    // Recursive DFS with information from upper layers
    const result = [];
    (function dfs(node, target, ancestors) {
        if (node) {
            const newAncestors = ancestors.concat([node.val]);
            if (node.left) {
                dfs(node.left, target - node.val, newAncestors);
            }
            if (node.val === target && !node.left && !node.right) {
                result.push(newAncestors);
            }
            if (node.right) {
                dfs(node.right, target - node.val, newAncestors);
            }
        }
    })(root, targetSum, []);
    return result;
};
