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
 * @param {number[]} to_delete
 * @return {TreeNode[]}
 */
var delNodes = function(root, to_delete) { // Remove nodes, usually the parent result depends on child result, so we can make a preOrder tranverse
    const result = [];
    root = (function postOrder(node) {
        if (node) {
            node.left = postOrder(node.left);
            node.right = postOrder(node.right);
            if (to_delete.indexOf(node.val) > -1) {
                if (node.left) {
                    result.push(node.left);
                }
                if (node.right) {
                    result.push(node.right);
                }
                return null;
            }
            return node;
        }
        return null;
    })(root);
    if (root) {
        result.push(root);
    }
    return result;
};
