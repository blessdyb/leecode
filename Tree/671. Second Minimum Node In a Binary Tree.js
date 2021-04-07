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
var findSecondMinimumValue = function(root) {
    const nodes = [];
    const stack = [root];
    while(stack.length) {
        const node = stack.pop();
        if (node) {
            nodes.push(node.val);
            if (node.val === root.val) {
                stack.push(node.right);
                stack.push(node.left);
            }
        }
    }
    let [min, result] = [root.val, Infinity];
    nodes.forEach(i => {
        if (min < i && i < result) {
          result = i; 
        }
    });
    return result !== Infinity ? result : -1;
};
