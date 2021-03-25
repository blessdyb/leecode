/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {number[]} inorder
 * @param {number[]} postorder
 * @return {TreeNode}
 */
var buildTree = function(inorder, postorder) {
    if (inorder.length) {
        const rootValue = postorder[postorder.length - 1];
        const rootIndex = inorder.indexOf(rootValue);
        
        const leftTreeInOrder = inorder.slice(0, rootIndex);
        const rightTreeInOrder = inorder.slice(rootIndex + 1);
        const leftTreePostOrder = postorder.slice(0, leftTreeInOrder.length);
        const rightTreePostOrder = postorder.slice(leftTreeInOrder.length, postorder.length - 1);
        
        const node = new TreeNode(rootValue);
        node.left = buildTree(leftTreeInOrder, leftTreePostOrder);
        node.right = buildTree(rightTreeInOrder, rightTreePostOrder);
        return node;
    }
    return null;
};
