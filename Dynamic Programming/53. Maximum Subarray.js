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

var maxSumArray = function(nums) {  // DP, base case will be the first element. Then any given dp[i] depends on dp[i - 1], either we add the value
    const dp = [nums[0]];  // Since nums could be all negative, so the initial case we will use dp[0] = nums[0] instead of dp[1] = nums[0]
    for (let i = 1; i < nums.length; i++) {
        dp[i] = Math.max(dp[i - 1] + nums[i], nums[i]);
    }
    return Math.max(...dp);
};

var maxSumArray = function(nums) { // Devide and Conquer
    if (nums.length === 0) {
        return -Infinity;
    } else if (nums.length < 2) {
        return nums[0];   
    }
    let max = -Infinity;
    let middle = Math.floor(nums.length / 2);
    let leftMax = maxSumArray(nums.slice(0, middle));
    let rightMax = maxSumArray(nums.slice(middle));
    
    let middleToLeftMax = 0;
    let middleToRightMax = 0;
    let current = 0;
    for (let i = middle - 1; i >= 0; i--) {
        current += nums[i];
        middleToLeftMax = Math.max(middleToLeftMax, current);
    }
    current = 0;
    for (let i = middle + 1; i < nums.length; i++) {
        current += nums[i];
        middleToRightMax = Math.max(middleToRightMax, current);
    }
    return Math.max(leftMax, rightMax, nums[middle] + middleToLeftMax + middleToRightMax);
}
