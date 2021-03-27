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
 * @param {number} val
 * @param {number} depth
 * @return {TreeNode}
 */
var addOneRow = function(root, val, depth) {
    if (depth === 1) { // Handle exception
        root = new TreeNode(val, root);
    } else {
        let layer = 1;
        const stack = [root];
        let count = 1;
        let current = 0;
        while(stack.length) {
            const node = stack.shift();
            if (node) {
                current++;
                if (layer + 1 === depth) {
                    node.left = new TreeNode(val, node.left);
                    node.right = new TreeNode(val, null, node.right);
                } else {
                    if (node.left) {
                        stack.push(node.left);
                    }
                    if (node.right) {
                        stack.push(node.right);
                    }
                }
                if (layer === depth) {
                    break;
                }
                if (current === count) {
                    current = 0;
                    count = stack.length;
                    layer++;
                }
            }
        }
    }
    return root;
};
