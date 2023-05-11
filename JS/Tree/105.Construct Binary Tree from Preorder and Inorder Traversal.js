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
 * @param {number[]} inorder
 * @return {TreeNode}
 */
var buildTree = function(preorder, inorder) {  // Just need to use inorder to find the root node, then we can split the given two number list as left and right part. Recursively perform the action until there's no elements left in the number list.
    if (preorder.length) {
        const rootValue = preorder[0];
        const rootIndex = inorder.indexOf(rootValue);
        
        const leftTreeInOrder = inorder.slice(0, rootIndex);
        const rightTreeInOrder = inorder.slice(rootIndex + 1);
      
        const leftTreePreOrder = preorder.slice(1, leftTreeInOrder.length + 1);
        const rightTreePreOrder = preorder.slice(leftTreeInOrder.length + 1);
      
        const root = new TreeNode(rootValue);
        root.left = buildTree(leftTreePreOrder, leftTreeInOrder);
        root.right = buildTree(rightTreePreOrder, rightTreeInOrder);
        return root;    
    }
    return null;
};
