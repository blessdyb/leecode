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
    return (function(start, end) {
        const result = [];
        for(let i = start; i <= end; i++) {
            generateTrees(start, i - 1).forEach(left => {
                generateTrees(i + 1, end).forEach(right => {
                    const node = new TreeNode(i, left, right);
                    result.push(node);
                });
            });
        }
        return result.length ? result : [null];
    })(1, n);
};

var generateTrees = function(n) {
    const dp = [[null], [new TreeNode(1)]];
    
    function cloneNode(node, offset) {
        if (node) {
            const newNode = new TreeNode(node.val + offset);
            newNode.left = cloneNode(node.left, offset);
            newNode.right = cloneNode(node.right, offset);
            return newNode;
        }
        return node;
    }
    
    
    for(let i = 2; i <= n; i++) {
        dp[i] = [];
        for(let j = 0; j < i; j++) {
            dp[j].forEach(left => {
                dp[i - j - 1].forEach(right => {
                    dp[i].push(new TreeNode(j + 1, left, cloneNode(right, j + 1)));  
                });
            });
        }
    }
    return dp[n];
}

