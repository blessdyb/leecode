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
var lcaDeepestLeaves = function(root) {
    let max = -1;
    let leaves = [];
    const stack = [[root, 0, []]];
    while(stack.length) {
        const [n, depth, ancestors] = stack.pop();
        if (n) {
            const newAncestors = ancestors.concat(n);
            if (!n.left && !n.right) {
                if (depth > max) {
                    leaves = [newAncestors];
                    max = depth;
                } else if (depth === max) {
                    leaves.push(newAncestors);
                }
            }
            if (n.left) {
                stack.push([n.left, depth+1, newAncestors]);
            }
            if (n.right) {
                stack.push([n.right, depth+1, newAncestors]);
            }
        }
    }
    if (leaves.length) {
        let cursor = 0;
        while(cursor <= max) {
            for (let i = 1; i < leaves.length; i++) {
                if (leaves[i][cursor].val !== leaves[0][cursor].val) {
                    return leaves[0][cursor - 1];
                }
            }
            cursor++;
        }
    }
    return leaves[0][max];
};

