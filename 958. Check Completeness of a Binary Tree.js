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
 * @return {boolean}
 */
var isCompleteTree = function(root) {  // Complete tree means for each layer, all node's previous node is not not null
    let nodeIsMissing = false;
    const stack = [root];
    let count = 1;
    let current = 0;
    while(stack.length) {
        const n = stack.shift();
        current++;
        if (n) {
            if (nodeIsMissing) {
                return false;
            }
            stack.push(n.left);    // So here we push all sub nodes and check if it's null when traversing next layer
            stack.push(n.right);
            if (current === count) {
                count = stack.length;
                current = 0;
                nodeIsMissing = false;
            }
        } else {
            nodeIsMissing = true;
        }
    }
    return true;
};
