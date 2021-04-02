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
var widthOfBinaryTree = function(root) {
// To calculate the layer width, the simplist way is to traverse the tree layer by layer and assign an index to each node. Then calculate the width. Since the 
// Index grows exponentially layer by layer, to avoid this issue, we can always use the left most value as an offset.
    let max = 0;
    let left = 0;
    let right = 0;
    let count = 1;
    let current = 0;
    const stack = [[root, 1]];
    while(stack.length) {
        const [node, index] = stack.shift();
        if (node) {
            if (current === 0) {
                left = index;
            }
            current++;
            if (node.left) {
                stack.push([node.left, index * 2 - left]);
            }
            if (node.right) {
                stack.push([node.right, index * 2 + 1 - left]);
            }
            if (current === count) {
                right = index;
                count = stack.length;
                current = 0;
                console.log(left, right);
                max = Math.max(right - left + 1, max);
            }
        }
    }
    return max;
};
