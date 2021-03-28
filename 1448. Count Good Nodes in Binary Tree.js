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
var goodNodes = function(root) {
    let result = 0;
    const stack = [[root, root.val]];
    while(stack.length) {
        let [node, maxAncestor] = stack.shift();
        if (node) {
            if (node.val >= maxAncestor) {
                result++;
                maxAncestor = node.val;
            }
            if (node.left) {
                stack.push([node.left, maxAncestor])
            }
            if (node.right) {
                stack.push([node.right, maxAncestor])
            }
        }
    }
    return result;
};
