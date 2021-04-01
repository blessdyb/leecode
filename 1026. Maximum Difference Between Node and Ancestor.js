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
var maxAncestorDiff = function(root) {
    let max = -Infinity;
    const stack = [[root, []]];
    while(stack.length) {
        const [node, ancestors] = stack.pop();
        if (node) {
            max = Math.max(...ancestors.map(i => Math.abs(i - node.val)).concat([max]));
            const newAncestors = ancestors.concat([node.val]);
            if (node.left) {
                stack.push([node.left, newAncestors]);
            }
            if (node.right) {
                stack.push([node.right, newAncestors]);
            }
        }
    }
    return max;
};
