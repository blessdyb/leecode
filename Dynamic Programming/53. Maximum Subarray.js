/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function(nums) { // Brute-force
    let result = -Infinity;
    for(let i = 0; i < nums.length; i++) {
        let temp = 0;
        for(let j = i; j < nums.length; j++) {
            temp += nums[j];
            result = Math.max(result, temp);
        }
    }
    return result;
};
