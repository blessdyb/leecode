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

var lowestCommonAncestor = function(root, p, q) {
    const stack = [root];
    const low = p.val > q.val ? q : p;
    const high = p.val > q.val ? p : q;
    while(stack.length) {
        const node = stack.shift();
        if (node.val === low.val || node.val === high.val) {
            return node;
        }
        if (node.val < low.val) {
            stack.push(node.right);
        } else if (node.val > high.val) {
            stack.push(node.left);
        } else {
            return node;
        }
    }
};
