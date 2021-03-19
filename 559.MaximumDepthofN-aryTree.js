/**
 * // Definition for a Node.
 * function Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {Node} root
 * @return {number}
 */
var maxDepth = function(root) {
    if (root) {
        const childs = root.children.map(c => maxDepth(c));
        return 1 + (childs.length ? Math.max(...childs) : 0);
    }
    return 0;
};
