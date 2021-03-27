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
 * @return {void} Do not return anything, modify root in-place instead.
 */
var flatten = function(root) { // Recursively flatten subleft and subright trees, and then concat root and the flatten sub trees
    if (root) {
        const flattenLeftTree = flatten(root.left);
        const flattenRightTree = flatten(root.right);
        root.left = null;
        if (flattenLeftTree) {
            let previous = flattenLeftTree.right;
            while(previous.right) {       // Note, we need to get the most right leaf node, and then set its right to the root node's flatten right subtree
                previous = previous.right; 
            }
            previous.right = flattenRightTree;
            root.right = flattenLeftTree;
        } else {
            root.right = flattenRightTree;
        }
    }
    return root;
};

var flatten = function(root) {
/**
 * The procedure can be DFS + Morris-like traversal
 * a. Get left subtree's right most leaf, set it to right subtree
 * b. Set right subtree to current left subtree and set left subtree to null
 * c. Continue this process until we reach the right most leaf
 */
    let current = root;
    let previous;
    while(current) {
        if (current.left) {
            previous = current.left;
            while(previous.right) {
                previous = previous.right;   
            }
            previous.right = current.right;
            current.right = current.left;
            current.left = null;
        } else {
            current = current.right;   
        }
    }
    return root;
};
