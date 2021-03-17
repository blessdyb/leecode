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
var minDepth = function(root) {
    if (root) {
        if (root.left && root.right) {
            return 1 + Math.min(minDepth(root.left), minDepth(root.right));
        } else if (root.left) {
            return 1 + minDepth(root.left);
        } else {
            return 1 + minDepth(root.right);
        }
    }
    return 0;
};

var minDepth = function(root) {
    let depth = 0;
    if (root) {
        const stack = [root];
        depth = 1;
        let count = 1;
        let cursor = 0;
        while(stack.length) {
            const node = stack.shift();
            cursor++;
            if (!node.left && !node.right) {
                break;   
            }
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
};
