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
 */
var BSTIterator = function(root) {
    this.current = root;
    this.stack = [];
};

/**
 * @return {number}
 */
BSTIterator.prototype.next = function() {  // Just the inOrder traversal stack version
    while(this.current) {
        this.stack.push(this.current);
        this.current = this.current.left;
    }
    const node = this.stack.pop();
    this.current = node.right;
    return node.val;
};

/**
 * @return {boolean}
 */
BSTIterator.prototype.hasNext = function() {
    return this.current || this.stack.length
};

/** 
 * Your BSTIterator object will be instantiated and called as such:
 * var obj = new BSTIterator(root)
 * var param_1 = obj.next()
 * var param_2 = obj.hasNext()
 */

// Brute-force solution will be creating a stack and pushing all nodes inside by inorder traversing, since we can't easily control recursive. So we can consider inOrder traverse without recursive
function inOrderTraverse(root) {
    const stack = [];
    let current = root;
    while(true) {
        if (current) {
            stack.push(current);
            current = current.left;
        } else if (stack.length) {
            const node = stack.pop();
            // Get the node as inOrder sequence
            current = node.right;
        } else {
            break;
        }
    }
}
