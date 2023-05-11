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
 * @return {number[]}
 */
var findFrequentTreeSum = function(root) {
    const map = {};
    let maxValueCount = 0;
    (function getFrequent(node) {
        if (node) {
            const value = node.val + getFrequent(node.left) + getFrequent(node.right);
            if (map[value] === undefined) {
                map[value] = 0;
            }
            map[value]++;
            if (map[value] > maxValueCount) {
                maxValueCount = map[value];
            }
            return value;
        }
        return 0;
    })(root);
    return Object.keys(map).filter(i => map[i] === maxValueCount);
};
