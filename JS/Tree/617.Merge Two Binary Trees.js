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
 * @return {TreeNode}
 */
var mergeTrees = function(root1, root2) {
    if (root1 && root2) {
        root1.val = root1.val + root2.val;
        root1.left = mergeTrees(root1.left, root2.left);
        root1.right = mergeTrees(root1.right, root2.right);
    }
    return root1 || root2;
};

var mergeTrees = function(root1, root2) {
    if (root1 && root2) { // Make sure we cover the edge case, since we need to create a new TreeNode if it's null. At the beginning, if root1 is null, we need to catch this case.
        const stack = [[root1, root2]];
        while(stack.length) {
            const [n1, n2] = stack.shift();
            n1.val += n2;
            if (n2.left) {
                if (!n1.left) {
                    n1.left = new TreeNode(0);   
                }
                stack.push([n1.left, n2.left]);
            }
            if (n2.right) {
                if (!n1.right) {
                    n1.right = new TreeNode(0);   
                }
                stack.push([n1.right, n2.right]);
            }
        }
    }
    return root1 || root2;
};
