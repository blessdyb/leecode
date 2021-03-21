/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} t
 * @return {string}
 */
var tree2str = function(t) {
    if (t) {
        let child = '';
        if (!t.left && !t.right) {
            child = '';   
        } else if (t.left && t.right) {
            child = `(${tree2str(t.left})(${tree2str(t.right})`;
        } else if (t.left) {
            child = `(${tree2str(t.left})`;
        } else {
            child = `()(${tree2str(t.left})`;
        }
        return `${t.val}${child}`;
    }
    return '';
};
// Since for each node, we need to wrap all it's node into parenthesis. It means we need to start with the parent node and end with the parent node. The traiditional
// tranverse iteration will only visit the parent node once since we will pop it. To make sure we can visit all nodes twice, we don't pop out the node at the first time,
// just reference it from the stack top and save it to another visited list. 
var tree2str = function(t) {
    if (!t) {
        return '';   
    }
    const stack = [t];
    const visited = [];
    let result = '';
    while (stack.length) {
        const n = stack[stack.length - 1];
        if (n) {
            if (visited.indexOf(n) > -1) {
                stack.pop();
                result += ')';
            } else {
                visited.push(n);
                result += '(' + n.val;
                if (!n.left && n.right) {
                    result += '()';
                }
                if (n.right) {
                    stack.push(n.right);
                }
                if (n.left) {
                    stack.push(n.left);
                }
            }
        }
    }
    
    return result.slice(1, result.length - 1);
};
