/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {number[]} preorder
 * @return {TreeNode}
 */
var bstFromPreorder = function(preorder) {
    if (preorder.length) {
        const node = new TreeNode(preorder[0]);
        const leftTreeNodes = preorder.filter(i => i < preorder[0]);  // Find the left subtree nodes, recursively build the tree
        node.left = bstFromPreorder(leftTreeNodes);
        node.right = bstFromPreorder(preorder.slice(leftTreeNodes.length + 1));
        return node;
    }
    return null;
};

var bstFromPreorder = function(preorder) {
    const root = new TreeNode(preorder[0]);
    let current = 1;
    const len = preorder.length;
    const stack = [root];
    while(current < len) {
        const node = new TreeNode(preorder[current]);
        let top = stack[stack.length - 1];
        // Ensure the nodes in stack are always bigger than current node
        // This can help us to pop the node out when we need to set its right tree
        while (stack.length && stack[stack.length - 1].val < node.val) {
            top = stack.pop();   
        }
        if (top.val < node.val) {
            top.right = node;   
        } else {
            top.left = node;
        }
        stack.push(node);
        current++;   
    }
    return root;
};
