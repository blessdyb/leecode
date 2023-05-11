/**
 * // Definition for a Node.
 * function Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {Node} root
 * @return {number[][]}
 */
var levelOrder = function(root) {
    const result = [];
    const stack = [root];
    let count = 1;
    let current = 0;
    let depth = 0;
    while(stack.length) {
        const node = stack.shift();
        if (node) {
            if (current === 0) {
                result.push([]);
            }
            current++;
            result[depth].push(node.val);
            node.children.forEach(i => stack.push(i));
            if (current === count) {
                current = 0;
                count = stack.length;
                depth++;
            }
        }
    }
    return result;
};
