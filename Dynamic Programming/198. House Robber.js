/**
 * @param {number[]} nums
 * @return {number}
 */
var rob = function(nums) {
    const robs = [0, nums[0]];
    const notRobs = [0, 0];
    const n = nums.length;
    for(let i = 2; i <= n; i++) {
        robs[i] = nums[i - 1] + Math.max(robs[i - 2], notRobs[i - 2]);
        notRobs[i] = Math.max(notRobs[i - 1], robs[i - 1]);
    }
    return Math.max(robs[n], notRobs[n]);
};
