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

var maxAncestorDiff = function(root) { // To get the max diff between node and its ancestors, we only need to compare the current node with tha max/min values from it's ancestors instead of all.
    let result = -Infinity;
    const stack = [[root, root.val, root.val]];
    while(stack.length) {
        let [node, min, max] = stack.pop();
        if (node) {
            result = Math.max(Math.abs(node.val - min), Math.abs(node.val - max), result);
            min = min > node.val ? node.val : min;
            max = max < node.val ? node.val : max;
            if (node.left) {
                stack.push([node.left, min, max]);
            }
            if (node.right) {
                stack.push([node.right, min, max]);
            }
        }
    }
    return result;
};
