/**
 * @param {number[]} nums
 * @return {number}
 */
var waysToMakeFair = function(nums) {
    const n = nums.length;
    const oddSum = [0];
    const evenSum = [0];
    for(let i = 0; i < n; i++) {
        oddSum[i + 1] = oddSum[i] + (i % 2 ? nums[i] : 0);
        evenSum[i + 1] = evenSum[i] + (i % 2 ? 0 : nums[i]);
    }
    let count = 0;
    for(let i = 0; i < n; i++) {
        const newOdd = oddSum[i] + (evenSum[n] - evenSum[i + 1]);
        const newEven = evenSum[i] + (oddSum[n] - oddSum[i + 1]);
        if (newOdd === newEven) {
            count++; 
        }
    }
    return count;
};  
