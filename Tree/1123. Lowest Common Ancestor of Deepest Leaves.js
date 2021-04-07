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
    let result = null;
    let maxDepth = 0;
    (function lca(node, depth) {
        maxDepth = Math.max(maxDepth, depth);
        if (node) {
            const left = lca(node.left, depth + 1);
            const right = lca(node.right, depth + 1);
            if (left === right && left === maxDepth) {
                result = node;   
            }
            return Math.max(left, right);
        }
        return depth;
    })(root, 0);
    return result;
}


var lcaDeepestLeaves = function(root) {  // DFS
    return (function lca(node) {
        if (node) {
            const [depth1, lca1] = lca(node.left);
            const [depth2, lca2] = lca(node.right);
            if (depth1 > depth2) {
                return [depth1, lca1];   
            } else if (depth1 < depth2) {
                return [depth2, lca2];   
            } else {
                return [depth1, node];   
            }
        }
        return [0, null];
    })(root)[1];
};

var lcaDeepestLeaves = function(root) { // BFS + Stack
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

