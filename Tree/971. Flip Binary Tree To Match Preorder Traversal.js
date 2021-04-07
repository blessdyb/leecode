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
 * @param {number[]} voyage
 * @return {number[]}
 */

// If a binary tree's all node values are unique. We can make a canonical traverse with preOrder DFS. And the canonical traverse will result a unique tree. 
// Similar here, we can use voyage as a given canonical-like traverse base and flip the tree's node on the fly.
var flipMatchVoyage = function(root, voyage) {
    let result = [];
    let cursor = 0;
    function dfs(node) {
        if(node) {
            if (node.val !== voyage[0]) {
                result = [-1];
                return;
            }
            cursor++;
            if (cursor < voyage.length && node.left && node.left.val !== voyage[cursor]) {
                result.push(node.val);
                dfs(node.right);
                dfs(node.left);
            } else {
                dfs(node.left);
                dfs(node.right);
            }
        }
    }
    dfs(root);
    return result.indexOf(-1) > -1 ? [-1] : result;
};
