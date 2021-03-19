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
var maxDepth = function(node) {
    if (node) {
        return 1 + Math.max(maxDepth(node.left), maxDepth(node.right));
    }
    return 0;
};

var maxDepth = function(node) {
    const stack = [node];
    let count = 1;
    let cursor = 0;
    let depth = 0;
    while(stack.length) {
        const node = stack.shift();
        if (node) {
            if (node.left) {
                stack.push(node.left);
            }
            if (node.right) {
                stack.push(node.right);
            }
            cursor++;
            if (cursor === count) {
                count = stack.length;
                cursor = 0;
                depth++;
            }
        }
    }
    return depth;   
}
// The diameter of a binary tree can be either
// 1. it's the diameter of the left sub tree if not passing through root node [fallback to recursive]
// 2. it's the diameter of the right sub tree if not passing through root node [fallback to recursive]
// 3. the sum of the max length of both of the root's child nodes [maxth depth of a binary tree, either BFS or recursive max + 1]
var diameterOfBinaryTree = function(root) {
    if (root) {
        return Math.max(diameterOfBinaryTree(root.left), diameterOfBinaryTree(root.right), maxDepth(root.left) + maxDepth(root.right));
    }
    return 0;
};
