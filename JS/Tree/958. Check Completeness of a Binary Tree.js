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
 * @return {boolean}
 */
var isCompleteTree = function(root) {  // Complete tree means for each layer, all node's previous node is not not null
    let nodeIsMissing = false;
    const stack = [root];
    let count = 1;
    let current = 0;
    while(stack.length) {
        const n = stack.shift();
        current++;
        if (n) {
            if (nodeIsMissing) {
                return false;
            }
            stack.push(n.left);    // So here we push all sub nodes and check if it's null when traversing next layer
            stack.push(n.right);
            if (current === count) {
                count = stack.length;
                current = 0;
                nodeIsMissing = false;
            }
        } else {
            nodeIsMissing = true;
        }
    }
    return true;
};

var isCompleteTree = function(root) {  // Complete tree means when we traverse the tree layber by layer, we should get a continus value list. So if we give each node index, we just need to compare the index of current node with previous one to make sure the offset is 1.
    const stack = [[root, 1]];
    const result = [];
    while(stack.length) {
        const [n, index] = stack.shift();
        if (n) {
            if (result.length && (result[result.length - 1] + 1) !== index) {
                return false;
            }
            result.push(index);
            if (n.left) {
                stack.push([n.left, index * 2]);
            }
            if (n.right) {
                stack.push([n.right, index * 2 + 1]);
            }
        }
    }
    return true;
};
