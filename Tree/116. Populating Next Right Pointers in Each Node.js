/**
 * // Definition for a Node.
 * function Node(val, left, right, next) {
 *    this.val = val === undefined ? null : val;
 *    this.left = left === undefined ? null : left;
 *    this.right = right === undefined ? null : right;
 *    this.next = next === undefined ? null : next;
 * };
 */

/**
 * @param {Node} root
 * @return {Node}
 */
var connect = function(root) {       // Recursive
    if (root) {
        (function c(left, right) {
            if (left && right) {
                left.next = right;
                c(left.left, left.right)
                c(left.right, right.left)
                c(right.left, right.right)
            }
        })(root.left, root.right);
    }
    return root;
};

var connect = function(root) { // BFS + stack
    const stack = [root];
    let count = 1;
    let current = 0;
    let previous;
    while(stack.length) {
        const node = stack.shift();
        current++;
        if (node) {
            if (previous) {
                previous.next = node;
            }
            previous = node;
            if (node.left) {
                stack.push(node.left);
            }
            if (node.right) {
                stack.push(node.right);
            }
            if (current === count) {
                current = 0;
                previous = null;
                count = stack.length;
            }
        }
    }
    return root;
};

var connect = function(root) { // Assume the upper layer's next link has been built, for the lower layer, we can benefit from the next pointer.
    let previous = root;
    while(previous && previous.left) {
        let current = previous;
        while(current) {
            current.left.next = current.right;
            if (current.next) {
                current.right.next = current.next.left;   
            }
            current = current.next;
        }
        previous = previous.left;   
    }
    return root;
};
