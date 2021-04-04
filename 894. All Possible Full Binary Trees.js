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
var allPossibleFBT = function(n) {
    if (n % 2 === 0) {
        return [];
    } else if (n === 1) {
        return [new TreeNode(0)];
    } else {
        const result = [];
        for (let i = 1; i < n; i+=2) {
             allPossibleFBT(i).forEach(j => { // left tree 1 ~ n - 2
                  allPossibleFBT(n - i - 1).forEach(k => { //right tree 1 ~ n - 2
                      const root = new TreeNode(0);
                      root.left = j;
                      root.right = k;
                      result.push(root);
                  });
             });
        }
        return result;
    }
};
