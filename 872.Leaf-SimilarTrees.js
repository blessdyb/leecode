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
var leafSimilar = function(root1, root2) {
    function getLeavesDFS(node) {
        let leaves = '';
        const stack = [];
        let current = node;
        while(true) {
            if (current) {
                stack.push(current);
                current = current.left;
            } else if (stack.length) {
                const node = stack.pop();
                if (!node.left && !node.right) {
                    leaves += ',' + node.val; // In case it's number string, we split the number with comma
                }
                current = node.right;
            } else {
                break;
            }
        }
        return leaves;
    }
    return getLeavesDFS(root1) === getLeavesDFS(root2);
};
