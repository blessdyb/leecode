/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {number[]} pre
 * @param {number[]} post
 * @return {TreeNode}
 */
var constructFromPrePost = function(pre, post) { // Recursive build tree
    if (pre.length) {
        const root = new TreeNode(pre[0]);
        if (pre.length > 1) {
            const leftRootNode = pre[1];
            const leftRootNodeIndex = post.indexOf(leftRootNode);
            root.left = constructFromPrePost(pre.slice(1, leftRootNodeIndex + 2), post.slice(0, leftRootNodeIndex + 1));
            root.right = constructFromPrePost(pre.slice(leftRootNodeIndex + 2), post.slice(leftRootNodeIndex + 1, post.length -1));
        }
        return root;
    }
    return null;
};
