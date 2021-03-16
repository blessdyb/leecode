/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {boolean}
 */
var isSameTree = function(p, q) {
    if (!p && !q) {
        return true;
    }
    if (!p|| !q) {
        return false;
    }
    const stack = [[p, q]];
    while(stack.length) {
        const temp = stack.shift()
        if (temp[0].val !== temp[1].val) {
            return false;
        }
        if ((temp[0].left && !temp[1].left) || (!temp[0].left && temp[1].left)) {
            return false;
        }
        if ((temp[0].right && !temp[1].right) || (!temp[0].right && temp[1].right)) {
            return false;
        }
        if (temp[0].left && temp[1].left) {
            stack.push([temp[0].left, temp[1].left])
        }
        if (temp[1].right && temp[1].right) {
            stack.push([temp[0].right, temp[1].right])
        }
    }
    return true;
};

var isSameTree = function(p, q) {
    if (!p && !q) {
        return true
    } else if (!p || !q) {
        return false
    } else if (p.val !== q.val) {
        return false
    } else {
        return isSameTree(p.left, q.left) && isSameTree(p.right, q.right)
    }
};
