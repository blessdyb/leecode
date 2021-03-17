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
 * @return {boolean}
 */
var hasPathSum = function(root, targetSum) {
    if (!root) {
        return false;
    }
    if (!root.left && !root.right) {
        return root.val === targetSum;
    }
    return hasPathSum(root.left, targetSum - root.val) || hasPathSum(root.right, targetSum - root.val);
};

var hasPathSum = function(root, targetSum) {
    if (!root) {
        return false;
    }
    const stack = [[root, targetSum]];
    while(stack.length) {
        const [node, num] = stack.pop();
        const diff = num - node.val;
        if (!node.left && !node.right && diff === 0) {
            return true;   
        }
        if (node.left) {
            stack.push([node.left, diff]);   
        }
        if (node.right) {
            stack.push([node.right, diff]);   
        }
    }
    return false;
};
