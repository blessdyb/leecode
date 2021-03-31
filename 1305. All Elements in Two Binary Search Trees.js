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
 * @return {number[]}
 */
var getAllElements = function(root1, root2) {
    const stack1 = [];
    const stack2 = [];
    let current1 = root1;
    let current2 = root2;
    while(true) {
        if (current1) {
            stack1.push(current1);
            current1 = current1.left;
        } else if (current2) {
            stack2.push(current2);
            current2 = current2.left;
        } else if (stack1.length && stack2.length) {
            const node1 = stack1[stack1.length - 1];
            const node2 = stack2[stack2.length - 1];
            if (node1.val > node2.val) {
                result.push(node2.val);
                stack2.pop();
                current2 = node2.right;
            } else {
                result.push(node1.val);
                stack1.pop();
                current1 = node1.right;
            }
        } else if (stack1.length) {
            const node = stack1.pop();
            result.push(node.val);
            current1 = node.right;
        } else if (stack2.length) {
            const node = stack2.pop();
            result.push(node.val);
            current2 = node.right;
        } else {
            break; 
        }
    }
    return result;
};
