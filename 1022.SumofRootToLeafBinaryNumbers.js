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
var sumRootToLeaf = function(root) {
    const stack = [[root, '']];
    const result = [];
    while(stack.length) {
        const [node, ancestor] = stack.shift();
        if (node) {
            const temp = ancestor + node.val;
            if (node.left) {
                stack.push([node.left, temp]);
            } 
            if (node.right) {
                stack.push([node.right, temp]);
            }
            if (!node.left && !node.right) {
                result.push(temp);
            }
        }
    }
    return result.reduce((acc, item) => acc + parseInt(item, 2), 0);
};
