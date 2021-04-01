/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {number[]} preorder
 * @return {TreeNode}
 */
var bstFromPreorder = function(preorder) {
    if (preorder.length) {
        const node = new TreeNode(preorder[0]);
        const leftTreeNodes = preorder.filter(i => i < preorder[0]);
        node.left = bstFromPreorder(leftTreeNodes);
        node.right = bstFromPreorder(preorder.slice(leftTreeNodes.length + 1));
        return node;
    }
    return null;
};
