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
    const stack = [[root, []]];
    let pPath, qPath;
    while(stack.length && (!pPath || !qPath)) {
        const [node, ancestor] = stack.shift();
        const newAncestor = ancestor.concat([node]);
        if (node.val === p.val) {
            pPath = newAncestor;
        }
        if (node.val === q.val) {
            qPath = newAncestor;
        }
        if (node.left) {
            stack.push([node.left, newAncestor]);
        }
        if (node.right) {
            stack.push([node.right, newAncestor]);
        }
    }
    let result = root;
    while(pPath.length && qPath.length) {
        const x = pPath.shift();
        const y = qPath.shift();
        if (x.val === y.val) {
            result = x;
        } else {
            break;
        }
    }
    return result;
};
