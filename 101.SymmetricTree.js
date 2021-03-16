/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {boolean}
 */
var isSymmetric = function(root) {
    var isChildSymmetric = function(left, right) {
        if (!left && !right) {
            return true;
        } else if (!left || !right) {
            return false;
        } else if (left.val !== right.val) {
            return false;
        } else {
            return isChildSymmetric(left.left, right.right) && isChildSymmetric(left.right, right.left)
        }
    }
    if (!root || (!root.left && !root.right)) {
        return true
    }
    return isChildSymmetric(root.left, root.right);
};

var isSymmetric = function(root) {
    if (!root || (!root.left && !root.right)) {
        return true
    }
    const stack = [[root.left, root.right]];
    while(stack.length) {
        const temp = stack.shift();
        if (!temp[0] && !temp[1]) {
            continue;
        } else if (!temp[0] || !temp[1]) {
            return false;
        } else if (temp[0].val !== temp[1].val) {
            return false;
        } else {
            stack.push([temp[0].left, temp[1].right]);
            stack.push([temp[0].right, temp[1].left])
        }
    }
    return true;
};
