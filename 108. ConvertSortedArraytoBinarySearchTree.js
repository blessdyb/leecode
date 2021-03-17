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

var sortedArrayToBST = function(nums) {
    if (nums.length) {
        const root = new TreeNode(nums[Math.floor(nums.length / 2)]);
        const stack = [[0, nums.length, root]];
        while(stack.length) {
            const [start, end, node] = stack.shift();
            const middle = Math.floor((start + end) / 2);
            
            if (middle > start) {
                const left = new TreeNode(nums[Math.floor((start + middle) / 2)]);
                node.left = left;
                stack.push([start, middle, left]);
            }
            
            if (middle < end - 1) {
                const right = new TreeNode(nums[Math.floor((middle + 1 + end) / 2)]);
                node.right = right;
                stack.push([middle + 1, end, right]);
            }
        }
        return root;
    }
    return null;
};
