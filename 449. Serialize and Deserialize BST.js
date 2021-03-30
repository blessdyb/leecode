/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */

/**
 * Encodes a tree to a single string.
 *
 * @param {TreeNode} root
 * @return {string}
 */
var serialize = function(root) {    // For BST, any node's left tree nodes always lower than its right nodes. With this in mind, we can easily inOrder traverse the tree
    const result = [];
    const stack = [root];
    while(stack.length) {
        const node = stack.pop();
        if (node) {
            result.push(node.val);
            if (node.right) {
                stack.push(node.right); 
            }
            if (node.left) {
                stack.push(node.left); 
            }
        }
    }
    return result.join(',');
};

/**
 * Decodes your encoded data to tree.
 *
 * @param {string} data
 * @return {TreeNode}
 */
// For an inOrder traversed BST node list, the first node will be it's root node. All values lower than the first element are adjacent elements in the array, so 
// we can recursively reconstruct the tree.
we can easily get the left nodes and right nodes.
var deserialize = function(data) {  
    if (data) {
        const values = data.split(',').map(item => parseInt(item, 10));
        return (function buildBST(list) {
            const len = list.length;
            if (len) {
                const node = new TreeNode(list[0]);
                let j = 1;
                while(j < len && list[j] < list[0]) {   // Split the data with root node to get left subtree and right subtree
                      j++;
                }
                node.left = buildBST(list.slice(1, j));
                node.right = buildBST(list.slice(j));
                return node;
            }
            return null;
        })(values);
    }
    return null;
};

/**
 * Your functions will be called as such:
 * deserialize(serialize(root));
 */
