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

var isValidBST = function(root) { // Morris traversal
    let current = root;
    let previous;
    let previousValue = Infinity;
    const isValid = () => previousValue !== Infinity ? previousValue < current.val : true;
    while(current) {
        if (!current.left) {
            if (!isValid()) {
                return false;    
            }
            previousValue = current.val;
            current = current.right;
        } else {
            previous = current.left;
            while (previous.right && previous.right !== current) { // Check the left subtree's right most leaf node of the current node, so we need to make sure we are not creating a dead loop
                previous = previous.right;
            }
            if (previous.right) {
                previous.right = null;
                if (!isValid()) {
                    return false;    
                }
                previousValue = current.val;
                current = current.right;
            } else {
                previous.right = current;
                current = current.left;
            }
        }
    }
    return true;
}
