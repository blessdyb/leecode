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

var findTargetSumWays = function(nums, target) {
    const n = nums.length;
    const sum = nums.reduce((acc, item) => acc + item, 0);
    const dp = new Array(n + 1).fill(0).map(item => new Array(2 * sum + 1).fill(0));
    dp[0][sum] = 1; // given target 0, with 0 nums, we have 1 way
    for(let i = 1; i <= n; i++) {
        for(let j = -sum; j <= sum; j++) {
            // State transform function:  dp[i][j + nums[i]] += dp[i][j] and dp[i][j - nums[i]] += dp[i][j] (consider it like a binary tree)
            if (j + nums[i - 1] <= sum) {
                dp[i][j + sum + nums[i - 1]] += dp[i - 1][j + sum];   
            }
            if (j - nums[i - 1] >= -sum) {
                dp[i][j + sum - nums[i - 1]] += dp[i - 1][j + sum];   
            }
        }
    }
    return target > sum || target < -sum ? 0 : dp[n][target + sum];
};

// Based on the state transform function, we can use a 1D DP
var findTargetSumWays = function(nums, target) {
    const n = nums.length;
    const sum = nums.reduce((acc, item) => acc + item, 0);
    let dp = new Array(2 * sum + 1).fill(0);
    dp[sum] = 1;
    for(let i = 1; i <= n; i++) {
        const temp = new Array(2 * sum + 1).fill(0);
        for(let j = -sum; j <= sum; j++) {
            if (j + nums[i - 1] <= sum) {
                temp[j + sum + nums[i - 1]] += dp[j + sum];   
            }
            if (j - nums[i - 1] >= -sum) {
                temp[j + sum - nums[i - 1]] += dp[j + sum];   
            }
        }
        dp = temp;
    }
    return target > sum || target < -sum ? 0 : dp[target + sum];
};

var findTargetSumWays = function(nums, target) {  // convert to 0-1 knapsack problem
    const n = nums.length;
    const sum = nums.reduce((acc, item) => acc + item, 0);
    // N - P = target =>   2N = target + sum => N = (target + sum) / 2
    const targetSum = (sum + target) / 2;
    if ((sum + target) % 2 === 1) {
        return 0;   
    }
    const dp = new Array(n + 1).fill(0).map(item => new Array(targetSum + 1).fill(0));
    dp[0][0] = 1;
    for(let i = 1; i <= n; i++) {
        for(let j = 0; j <= targetSum; j++) {
            dp[i][j] = dp[i - 1][j] + (j - nums[i - 1] >= 0 ? dp[i - 1][j - nums[i - 1]] : 0);   
        }
    }
    return dp[n][targetSum];
};

var findTargetSumWays = function(nums, target) {
    const n = nums.length;
    const sum = nums.reduce((acc, item) => acc + item, 0);
    const targetSum = (sum + target) / 2;
    if ((sum + target) % 2) {
        return 0;   
    }
    let dp = new Array(targetSum + 1).fill(0);
    dp[0] = 1;
    for(let i = 1; i <=n; i++) {
        const temp = dp.slice(0);
        for(let j = 0; j <= targetSum - nums[i - 1]; j++) {
            temp[j + nums[i - 1] += dp[j]; // push with rolling 1D DP
        }
        dp = temp;
    }
    return dp[targetSum];
}

var findTargetSumWays = function(nums, target) {
    const n = nums.length;
    const sum = nums.reduce((acc, item) => acc + item, 0);
    const targetSum = (sum + target) / 2;
    if ((sum + target) % 2) {
        return 0;   
    }
    let dp = new Array(targetSum + 1).fill(0);
    dp[0] = 1;
    for(let i = 1; i <=n; i++) {
        for(let j = targetSum; j >= nums[i - 1]; j--) {
            dp[j] += dp[j - nums[i - 1]]; // pull 1D DP
        }
    }
    return dp[targetSum];
}
