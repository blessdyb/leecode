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
var serialize = function(root) {
    if (root) {
        const result = [];
        const stack = [root];
        while(stack.length) {  // BSF layer iteration
            const node = stack.shift();
            if (node) {
                result.push(node.val);
                stack.push(node.left); 
                stack.push(node.right);
            } else {
                result.push('#'); 
            }
        }
        return result.join(',');
    }
    return '';
};

/**
 * Decodes your encoded data to tree.
 *
 * @param {string} data
 * @return {TreeNode}
 */
var deserialize = function(data) {
    if (data) {
        const values = data.split(',').map(item => {
            if (item === '#') {
                return null; 
            }
            return parseInt(item, 10);
        });
        const root = new TreeNode(values[0]);
        let cursor = 0;
        const stack = [root];
        while(stack.length) {            // we have the value list which we can access the elements with index. So when we constructing the tree, we just need to 
                                         // hold a cursor when inserting nodes to the tree
            const node = stack.shift();  // Since we are constructing the tree, no need to verify the node
            cursor++;
            if (values[cursor] !== null) {
                node.left = new TreeNode(values[cursor]);
                stack.push(node.left);
            }
            cursor++;
            if (values[cursor] !== null) {
                node.right = new TreeNode(values[cursor]);
                stack.push(node.right);
            }
        }
        return root;
    }
    return null;
};

/**
 * Your functions will be called as such:
 * deserialize(serialize(root));
 */
