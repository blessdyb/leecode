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
 * @return {number}
 */
// BST means it's a sorted array when tranversing with DFS, min absolute diff only happens in adjacent elements, so we just need to remember previous node when tranversing
var getMinimumDifference = function(root) {
    let min = Infinity;
    let prev;
    (function dfs(node){
        if (node) {
            dfs(node.left);
            if (prev) {
                min = Math.min(min, Math.abs(prev.val - node.val));   
            }
            prev = node;
            dfs(node.right);
        }
        return min;
    })(root);
    return min;
};

// Non-recursive DFS
var getMinimumDifference = function(root) {
    let min = Infinity;
    const stack = [];
    let current = root;
    let previous;
    while(true) {
        if (current) {
            stack.push(current);
            current = current.left;
        } else if (stack.length) {
            current = stack.pop();
            if (previous) {
                min = Math.min(min, Math.abs(previous.val - node.val));   
            }
            previous = current;
            current = current.right;
        } else {
            break;
        }
    }
    return min;
};
