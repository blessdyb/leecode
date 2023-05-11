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
 * @param {number} distance
 * @return {number}
 */
// Since we need to calculate the distance between any two leaf nodes, the paths will need to go through all parent nodes
// It means we can use postOrder DFS and populate all leaf nodes distance to it's LCA node.
var countPairs = function(root, distance) { 
    let result = 0;
    (function postOrderDFS(node) {
        if (node) {
            if (!node.left && !node.right) {  // Leaf node
                return [1]; 
            }
            const leftLeafNodes = postOrderDFS(node.left);
            const rightLeafNodes = postOrderDFS(node.right);
            leftLeafNodes.forEach(l => {
                rightLeafNodes.forEach(r => {
                    if (l + r <= distance) {
                        result++; 
                    }
                });
            });
            return leftLeafNodes.concat(rightLeafNodes).map(i => i + 1);
        }
        return [];
    })(root);
    return result;
};
