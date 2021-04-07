/**
 * @param {number[]} nums
 */
var NumArray = function(nums) {
    this.nums = nums;
};

/** 
 * @param {number} left 
 * @param {number} right
 * @return {number}
 */
NumArray.prototype.sumRange = function(left, right) {    // Brute force
    return this.nums.slice(left, right + 1).reduce((acc, item) => acc + item, 0);
};

/** 
 * Your NumArray object will be instantiated and called as such:
 * var obj = new NumArray(nums)
 * var param_1 = obj.sumRange(left,right)
 */



// In case sumRange gets called thousands time, we can cache the sum from 0 ~ i, so i ~ j = 0 ~ j - 0 ~ i;
var NumArray = function(nums) {
    this.nums = nums;
    this.sums = nums.reduce((acc, item, index) => { 
      acc.push(acc[index] + item);
      return acc;
    }, [0]);
};

NumArray.prototype.sumRange = function(left, right) {
    return this.sums[right + 1] - this.sums[left]; // Note the element is inclusive, so we need to have offset
};
