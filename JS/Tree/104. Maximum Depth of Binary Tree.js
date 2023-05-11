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
var maxDepth = function(root) {
    if (!root) {
        return 0;
    }
    return 1 + Math.max(maxDepth(root.left), maxDepth(root.right));
};

var maxDepth = function(root) {
    let depth = 0;
    if (root) {
        const stack = [root];
        let count = stack.length;
        let cursor = 0;
        while(stack.length) {
            const node = stack.shift();
            cursor++;
            if (node.left) {
                stack.push(node.left);
            }
            if (node.right) {
                stack.push(node.right);
            }
            if (cursor === count) {
                depth++;
                count = stack.length;
                cursor = 0;
            }
        }
    }
    return depth;
}
