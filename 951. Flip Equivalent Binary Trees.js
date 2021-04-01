/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root1
 * @param {TreeNode} root2
 * @return {boolean}
 */
var flipEquiv = function(root1, root2) {
    if (!root1 && !root2) {
        return true;
    } else if (root1 && root2) {
        if (root1.val !== root2.val) {
            return false;
        }
        return (flipEquiv(root1.left, root2.left) && flipEquiv(root1.right, root2.right)) || (flipEquiv(root1.left, root2.right) && flipEquiv(root1.right, root2.left))
    }
    return false;
};

// Since all nodes are uniqe, so we can swap the left/right subtrees in fly during preOrder DFS traversing by folliwing a rule (for example, left node always smaller than right node)
var flipEquiv = function(root1, root2) {
    function dfs(node, vals) {
        if (node) {
            vals.push(node.val);
            const leftVal = node.left ? node.left.val : -1;
            const rightVal = node.right ? node.right.val : -1;
            if (leftVal < rightVal) {
                dfs(node.left, vals);
                dfs(node.right, vals);
            } else {
                dfs(node.right, vals);
                dfs(node.left, vals);
            }
            vals.push(-2);
        }
    }
    
    const list1 = [];
    const list2 = [];
    dfs(root1, list1);
    dfs(root2, list2);
    return list1.join(',') === list2.join(',');
};
