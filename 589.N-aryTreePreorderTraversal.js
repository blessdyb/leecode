/**
 * // Definition for a Node.
 * function Node(val, children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {Node} root
 * @return {number[]}
 */
var preorder = function(root) {
    const stack = [root];
    const result = [];
    while(stack.length) {
        const node = stack.pop();
        if (node) {
            result.push(node.val);
            if (node.children) {
                node.children.reverse().forEach(item => stack.push(item));
            }
        }
    }
    return result;
};
