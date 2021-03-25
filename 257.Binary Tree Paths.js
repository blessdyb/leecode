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
 * @return {string[]}
 */
var binaryTreePaths = function(root) {
    const paths = [];
    const tranverse = function(node, ancestor) {
        const points = ancestor.concat([node.val]);
        if (!node.left && !node.right) {
            paths.push(points.join('->'));
        }
        if (node.left) {
            tranverse(node.left, points);   
        }
        if (node.right) {
            tranverse(node.right, points);
        }
    }
    tranverse(root, []);
    return paths;
};

var binaryTreePaths = function(root) {
    const paths = [];
    const stack = [[root, []]];
    while(stack.length) {
        const [node, ancestor] = stack.shift(); // Note: BFS we are always get the item from the bottom of the stack, layer by layer
        const updatedAncestor = ancestor.concat(node.val);
        if (!node.left && !node.right) {
            paths.push(updatedAncestor.join('->'));
        }
        if (node.left) { // tranverse layer from left to right, layer by layer
            stack.push([node.left, updatedAncestor]);
        }
        if (node.right) {
            stack.push([node.right, updatedAncestor])
        }
    }
    return paths;
};

var binaryTreePaths = function(root) {
    const paths = [];
    const stack = [[root, []]];
    while(stack.length) {
        const [node, ancestor] = stack.pop(); // Note: DFS we are always get the item from the top of the stack
        const updatedAncestor = ancestor.concat(node.val);
        if (!node.left && !node.right) {
            paths.push(updatedAncestor.join('->'));
        }
        if (node.right) { // always push the right node into the stack first for later usage, then push the left node to make sure the left node can get pop out from the top of the stack
            stack.push([node.right, updatedAncestor]);
        }
        if (node.left) {
            stack.push([node.left, updatedAncestor]);
        }
        
    }
    return paths;
};
