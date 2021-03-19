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

var maxDepth = function(root) {
    const stack = [root];
    let depth = 0;
    let count = 1;
    let cursor = 0;
    while(stack.length) {
        const node = stack.shift();
        if (node) {
            node.children.forEach(c => stack.push(c));
            cursor++;
            if (cursor === count) {
                cursor = 0;
                count = stack.length;
                depth++;
            }
        }
    }
    return depth;
 };
