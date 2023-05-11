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
 * @return {number[]}
 */
var preorderTraversal = function(root) {  // Recursive preOrder
    const result = [];
    (function preorder(node) {
        if (node) {
            result.push(node.val);
            preorder(node.left);
            preorder(node.right);
        }
    })(root);
    return result;
};

var preorderTraversal = function(root) { // Stack preOrder
    const result = [];
    const stack = [root];
    while(stack.length) {
        const node = stack.pop();
        if (node) {
            result.push(node.val);
            stack.push(node.right);
            stack.push(node.left);
        }
    }
    return result;
};

var preorderTraversal = function(root) { // Morris traversal
    const result = [];
    let current = root;
    let previous;
    while(current) {
        if (!current.left) {
            result.push(current.val);
            current = current.right;
        } else {
            previous = current.left;
            while(previous.right && previous.right !== current) {
                previous = previous.right;
            }
            if (!previous.right) {
                previous.right = current;
                result.push(current.val);
                current = current.left;
            } else {
                previous.right = null;
                current = current.right;
            }
        }
    }
    return result;
};
