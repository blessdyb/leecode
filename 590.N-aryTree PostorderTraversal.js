/**
 * // Definition for a Node.
 * function Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {Node} root
 * @return {number[]}
 */
var postorder = function(root) {
    const result =  [];
    (function dfs(node) {
        if (node) {
            if (node.children) {
                node.children.forEach(i => dfs(i));
            }
            result.push(node.val);
        }
    })(root);
    return result;
};
