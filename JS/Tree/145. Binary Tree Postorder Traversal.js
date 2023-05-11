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
var postorderTraversal = function(root) {  // Recursive
    const result = [];
    (function postorder(node) {
        if (node) {
            postorder(node.left);
            postorder(node.right);
            result.push(node.val);
        }
    })(root);
    return result;
};

var postorderTraversal = function(root) {  // Every time push the node twice so after we pop the node out, we can still have chance to visit it [Post Order]
    const result = [];
    const stack = [root, root];
    while(stack.length) {
        const node = stack.pop();
        if (node) {
            if (stack.length && node === stack[stack.length - 1]) { // If the current node equals to the one on top of the stack, it means we are traversing to left node
                stack.push(node.right);
                stack.push(node.right);
                stack.push(node.left);
                stack.push(node.left);
            } else {
                result.push(node.val);  // if the current node doesn't equal to the one on top of the stack, it means we are backing from the right subtree traversing
            }
        }
    }
    return result;
};
