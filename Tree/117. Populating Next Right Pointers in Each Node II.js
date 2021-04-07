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
var connect = function(root) {
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
            if (count === current) {
                count = stack.length;
                current = 0;
                previous = null;
            }
        }
    }
    return root;
};

var connect = function(root) {       // If we know how to find the start of each layer, then it's easier. Since we don't want to use stack, so the approach is to create a node and point to the first node of next layer.
    if (root) {
        let current = root;
        while(current) {
            let layerHead = new Node();  // User a layer head pointer to start from the under layer and point to the first node in the next layer
            let pointer = layerHead;
            while(current) {  
                if (current.left) {
                    pointer.next = current.left;   
                    pointer = pointer.next;
                }
                if (current.right) {
                    pointer.next = current.right;
                    pointer = pointer.next;
                }
                current = current.next;   // Horizontally move to the next node
            }
            current = layerHead.next;     // Move to the next layer
        }
    }
    return root;
};

