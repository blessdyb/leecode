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
