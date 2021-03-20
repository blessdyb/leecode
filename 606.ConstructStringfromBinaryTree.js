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
