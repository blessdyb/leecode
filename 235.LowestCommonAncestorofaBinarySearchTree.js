/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */

/**
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */
var lowestCommonAncestor = function(root, p, q) {
    if (root.val === p.val || root.val === q.val) {
        return root;
    }
    const low = p.val > q.val ? q : p;
    const high = p.val > q.val ? p : q;
    if (root.val < low.val) {
        return lowestCommonAncestor(root.right, low, high);
    } else if (root.val > high.val) {
        return lowestCommonAncestor(root.left, low, high);
    } else {
        return root;
    }
};
