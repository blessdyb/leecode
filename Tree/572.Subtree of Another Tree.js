/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} s
 * @param {TreeNode} t
 * @return {boolean}
 */
var isSubtree = function(s, t) {
    function isSameTree(a, b) {
        if (!a && !b) {
            return true;   
        } else if (!a || !b) {
            return false;   
        }
        return a.val === b.val && isSameTree(a.left, b.left) && isSameTree(a.right, b.right);
    }
    
    function tranverse(a, b) {
        return sameTree(a, b) || tranverse(a.left, b) || tranverse(a.right, b);
    }
    
    tranverse(s, t);
};

var isSubtree = function(s, t) {
    function treeToString(node, flag) {
        if (node) {
            return `#${node.val}${treeToString(node.left, '<')}${treeToString(node.right, '>')}`;
        }
        return flag;
    }
    const sStr = treeToString(s);
    const tStr = treeToString(t);
    return sStr.indexOf(tStr) > -1;
};
