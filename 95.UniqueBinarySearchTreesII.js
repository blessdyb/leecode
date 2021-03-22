/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {number} n
 * @return {TreeNode[]}
 */
var generateTrees = function(n) {
    return (function generate(start, end){
        const trees = [];
        for (let i = start; i <= end; i++) {
            generate(start, i - 1).forEach(j => {
               generate(i + 1, end).forEach(k => {
                  trees.push(new TreeNode(i, j, k)); 
               });
            });
        }
        return trees.length ? trees : [null];
    })(1, n);
};
