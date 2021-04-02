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
 * @param {number} low
 * @param {number} high
 * @return {TreeNode}
 */
var trimBST = function(root, low, high) {
    if (root) {
        if (root.val < low || root.val > high) {
            if (root.right) {
                let previous = root.right;
                while(previous.left) {
                    previous = previous.left;   
                }
                previous.left = root.left;
                return trimBST(root.right, low, high);
            }
            return trimBST(root.left, low, high);
        } else {
            root.left = trimBST(root.left, low, high);
            root.right = trimBST(root.right, low, high);
        }
    }
    return root;
};
