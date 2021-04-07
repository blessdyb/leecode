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

var postorder = function(root) {
    const stack = [root];
    const result = [];
    while(stack.length) {
        const node = stack.pop();
        if (node) {
            result.unshift(node.val);
            node.children.forEach(i => stack.push(i));
        }
    }
    return result;
};
