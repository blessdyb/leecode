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
 * @param {number} sum
 * @return {number}
 */
function numberOfPath(node, left) {
    if (node) {
        // Note: even we reached to the result, we need to continue check if the nodes sum under is 0, if it's 0, it means we get a longer path.
        return (node.val === left ? 1 : 0) + numberOfPath(node.left, left - node.val) + numberOfPath(node.right, left - node.val);
    }
    return 0;
}

var pathSum = function(root, sum) {  // The problem is asking for how many subPaths, so 
    if (root) {
        return numberOfPath(root, sum) + pathSum(root.left, sum) + pathSum(root.right, sum);
    }
    return 0;
};
