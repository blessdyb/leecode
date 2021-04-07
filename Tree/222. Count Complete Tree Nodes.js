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
var countNodes = function(root) {
    /**
     * We can easily traverse the tree with O(n) time complexity. To speed up, we can use the geometric progression formula 
     * to get the result if it's a perfect binary tree.
     * Sn = 1 + 2 + 4 + .... + Math.pow(2, n - 1)
     * Sn * 2 = 2 + 4 + 8 + ... + Math.pow(2, n - 1) + Math.pow(2, n)
     * Sn = Math.pow(2, n)  - 1 = (2 << n) - 1
     */
    if (!root) {    // Base case
        return 0;
    }
    let leftTreeHeight = 0;
    let rightTreeHeight = 0;
    let pointer = root;
    while(pointer.left) {
        pointer = pointer.left;
        leftTreeHeight++;
    }
    pointer = root;
    while(pointer.right) {
        pointer = pointer.right;
        rightTreeHeight++;
    }
    if (leftTreeHeight === rightTreeHeight) {  // If the left most leaf height === the right most leaf height, it's a full complete binary tree.
        return (2 << leftTreeHeight) - 1;
    } else {  // Recursive
        return countNodes(root.left) + countNodes(root.right) + 1;
    }
};
