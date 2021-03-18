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
var sumOfLeftLeaves = function(root) {
    if (root) {
        if (root.left && !root.left.left && !root.left.right) {
            return root.left.val + sumOfLeftLeaves(root.right);
        }
        return sumOfLeftLeaves(root.left) + sumOfLeftLeaves(root.right);
    }
    return 0;
};

var sumOfLeftLeaves = function(root) {
    let sum = 0;
    const stack = [root];
    while (stack.length) {
        const node = stack.pop();
        if (node) {
            if (node.left && !node.left.left && !node.left.right) {
                sum += node.left.val;   
            } else {
                stack.push(node.left);   
            }
            stack.push(node.right);
        }
    }
    return sum;
};


