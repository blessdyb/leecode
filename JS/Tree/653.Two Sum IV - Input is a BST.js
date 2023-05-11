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
 * @param {number} k
 * @return {boolean}
 */
var findTarget = function(root, k) { // Considering BST is a sorted array when tranversing with DFS (inOrderTranversing), so convert it to an sorted array then solve the 2Sum problem
    function convertBSTtoArray(node) {
        const stack = [];
        const list = [];
        let current = node;
        let cursor = 0;
        while(true) {
            if (current) {
                stack.push(current);
                current = current.left;
            } else if (stack.length) {
                const node = stack.pop();
                list.push(node.val);
                current = node.right;
            } else {
                break;
            }
        }
        return list;
    }
    const list = convertBSTtoArray(root);
    let i = 0;
    let j = 1;
    const len = list.length;
    while (i < j && j < len) {
        if (list[i] + list[j] === k) {
            return true;   
        }
        j++;
        if (j === len) {
            i++;
            j = i + 1;
        }
    }
    return false;
};
