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
var sumEvenGrandparent = function(root) {
    let result = 0;
    const stack = [[root, null]];
    while(stack.length) {
        const [node, parent] = stack.shift();
        if (node) {
            if (node.left) {
                stack.push([node.left, node]);
                if (parent && parent.val %2 === 0) {
                    result += node.left.val;
                }
            } 
            if (node.right) {
                stack.push([node.right, node]);
                if (parent && parent.val %2 === 0) {
                    result += node.right.val;
                }
            } 
        }
    }
    return result;
};
