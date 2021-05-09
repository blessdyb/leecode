/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var findTargetSumWays = function(nums, target) {
    let result = 0;
    const n = nums.length;
    function helper(i, t) {
        if (i === 0) {
            if (nums[0] === t) { // In case nums[i] is 0, we need to count it twice
                result++;
            }
            if(nums[0] === -t) {
                result++;
            }
        } else {
            helper(i - 1, t + nums[i]);
            helper(i - 1, t - nums[i]);
        }
    }
    helper(n - 1, target);
    return result;
};
