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
