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
 * @return {boolean}
 */
var isValidBST = function(root) { // Based on the BST defination, we can get a sorted array while traversing the tree by DFS. So just need to compare the adjacent element.
    const stack = [];
    let current = root;
    let previous;
    while(true) {
        if (current) {
            stack.push(current); 
            current = current.left;
        } else if (stack.length) {
            const node = stack.pop();
            if (previous && previous.val >= node.val) {
                return false; 
            }
            previous = node;
            current = node.right;
        } else {
            break;
        }
    }
    return true;
};
