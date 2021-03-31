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
// Each subtree's total depends on it's subtree, so we can traverse the tree with postOrder, this way we can easily get all subTree's sum in O(N)
var maxProduct = function(root) {
    let total = 0;
    let result = 0;
    function getNodesSum(node, checkResult) {
        if (node) {
            const leftTreeSum = getNodesSum(node.left, checkResult);
            const rightTreeSum = getNodesSum(node.right, checkResult);
            if (checkResult) {
                result = Math.max(result, leftTreeSum * (total - leftTreeSum), (total- rightTreeSum) * rightTreeSum);   // Cut left tree or right tree
            }
            return node.val + leftTreeSum + rightTreeSum;
        }
        return 0;
    }
    total = getNodesSum(root);
    getNodesSum(root, true);
    return result % (Math.pow(10, 9) + 7)
    
};
