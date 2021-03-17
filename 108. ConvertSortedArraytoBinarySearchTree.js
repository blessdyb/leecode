/**
 * Definition for a binary tree node.
 */
function TreeNode(val, left, right) {
    this.val = (val===undefined ? 0 : val)
    this.left = (left===undefined ? null : left)
    this.right = (right===undefined ? null : right)
}

/**
 * @param {number[]} nums
 * @return {TreeNode}
 */
var sortedArrayToBST = function(nums) {
    if (nums.length) {
        const middle = Math.floor(nums.length / 2);
        const left = middle > 0 ? sortedArrayToBST(nums.slice(0, middle)) : undefined;
        const right = middle < nums.length - 1 ? sortedArrayToBST(nums.slice(middle + 1)) : undefined;
        const node = new TreeNode(nums[middle], left, right);
        return node;
    }
    return null;
};
