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
 * @return {TreeNode}
 */
var bstToGst = function(root) {
    let sum = 0;
    let previous;
    let current = root;
    while(current) {
        if (!current.right) {
            sum += current.val;
            current.val = sum;
            previous = current;
            current = current.left;
        } else {
            previous = current.right;
            while(previous.left && previous.left !== current) {
                previous = previous.left;
            }
            if (previous.left) {
                previous.left = null;;
                sum += current.val;
                current.val = sum;
                current = current.left
            } else {
                previous.left = current;
                current = current.right;
            }
        }
    }
    return root;
};
