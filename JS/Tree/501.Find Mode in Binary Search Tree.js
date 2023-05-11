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
// DFS + BST can be considered as tranversing a sorted array. For BST with duplicated numbers, we have to use stack to tranverse since both left & right child nodes can have the same number as parent node.
var findMode = function(root) {
    let result = [];
    let current = 0;
    let count = 0; // Since the node val includes 0, we shouldn't initial it as 0 as a best practise. Here it doesn't matter since it's used to reset count to 0.
    let maxCount = 0;
    (function dfs(node) {
        if (node) {
            dfs(node.left);
            if (node.val != current) {
                current = node.val;
                count = 0;
            }
            count++;
            if (maxCount === count) {
                result.push(current);  
            } else if (maxCount < count) {
                result = [current];
                maxCount = current;
            }
            dfs(node.right);
        }
    })(root);
    return result;
};
// Tranverse first to know the result number, so we can save the result array space
var findMode = function(root) {
    let result;
    let current;
    let count = 0;
    let maxCount = 0;
    let resultNum = 0;
    
    function dfs(node) {
        if (node) {
            dfs(node.left);
            if (node.val !== current) {
                current = node.val;
                count = 0;
            }
            count++;
            if (count > maxCount) { // In the second round, this condition is always false
                maxCount = count;
                resultNum = 1;
            } else if (count === maxCount) {
                if (result) {
                    result[resultNum] = current; 
                }
                resultNum++;   
            }
            dfs(node.right);
        }
    }
    
    dfs(root); // To get maxCount and resultNum
    result = Array(resultNum);
    count = 0;
    resultNum = 0;
    dfs(root);
    return result;
}
