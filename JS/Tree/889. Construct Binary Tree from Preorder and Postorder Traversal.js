/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {number[]} pre
 * @param {number[]} post
 * @return {TreeNode}
 */
var constructFromPrePost = function(pre, post) { // Recursive build tree
    if (pre.length) {
        const root = new TreeNode(pre[0]);
        if (pre.length > 1) {
            const leftRootNode = pre[1];
            const leftRootNodeIndex = post.indexOf(leftRootNode);
            root.left = constructFromPrePost(pre.slice(1, leftRootNodeIndex + 2), post.slice(0, leftRootNodeIndex + 1));
            root.right = constructFromPrePost(pre.slice(leftRootNodeIndex + 2), post.slice(leftRootNodeIndex + 1, post.length -1));
        }
        return root;
    }
    return null;
};

var constructFromPrePost = function(pre, post) { // When preOrder and postOrder reach to a same node, it means we build the whole subtree.
    const stack = [new TreeNode(pre[0])];
    let cursor = 0;
    pre.slice(1).forEach(val => {
        const node = new TreeNode(val);
        while(stack[stack.length -1].val === post[cursor]) {
            stack.pop();
            cursor++;
        }
        if (stack[stack.length -1].left) {
            stack.left = node;   
        } else {
            stack.right = node;   
        }
        stack.push(node);
    });
    return stack[0];
};
