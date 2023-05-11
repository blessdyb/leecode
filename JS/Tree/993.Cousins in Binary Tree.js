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
 * @param {number} x
 * @param {number} y
 * @return {boolean}
 */
var isCousins = function(root, x, y) {
    function depthAndParent(node, val) {
        const stack = [[root, 0]];
        let count = 1;
        let cursor = 0;
        let depth = 0;
        while(stack.length) {
            const [node, parent] = stack.shift();
            if (node) {
                cursor++;
                if (node.val === val) {
                    return [depth, parent];
                }
                if (node.left) {
                    stack.push([node.left, node.val]);
                }
                if (node.right) {
                    stack.push([node.right, node.val]);
                }
                if (cursor === count) {
                    count = stack.length;
                    cursor = 0;
                    depth++;
                }
            }
        }
        return [-1, -1];
    }
    const xInfo = depthAndParent(root, x);
    const yInfo = depthAndParent(root, y);
    return xInfo[0] !== -1 && yInfo[0] !== -1 && xInfo[0] == yInfo[0] && xInfo[1] !== yInfo[1];
};
