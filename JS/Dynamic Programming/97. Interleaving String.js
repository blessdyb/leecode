/**
 * @param {string} s1
 * @param {string} s2
 * @param {string} s3
 * @return {boolean}
 */
var isInterleave = function(s1, s2, s3) {    // Recursive
    if (s1.length + s2.length !== s3.length) {
        return false; 
    } else if (s1.length === 0 && s2.length === 0 && s3.length === 0) {
        return true; 
    } else {
        let result = false;
        if (s1.length && s1[0] === s3[0]) {  // Note, we can't assume the intersect item sequence, so here we have to check both s1[0] === s3[0] and s2[0] === s3[0]
            result = result || isInterleave(s1.substr(1), s2, s3.substr(1)); 
        }
        if (s2.length && s2[0] === s3[0]) {
            result = result || isInterleave(s1, s2.substr(1), s3.substr(1)); 
        }
        return result;
    }
};

var isInterleave = function(s1, s2, s3) {   // DP
    if (s1.length + s2.length !== s3) {
        return false;   
    }
    const dp = [];
    for(let i = 0; i <= s1.length; i++) {
        dp[i] = [];
        for(let j = 0; j <= s2.length; j++) {
            if (i === 0 && j === 0) {
                dp[i][j] = true;   
            } else if (i === 0) {
                dp[i][j] = dp[i][j - 1] && s2[j - 1] === s3[i + j - 1];
            } else if (j === 0) {
                dp[i][j] = dp[i - 1][j] && s1[i - 1] === s3[i + j - 1];
            } else {
                dp[i][j] = (dp[i - 1][j] && s1[i - 1] === s3[i + j - 1]) || (dp[i][j - 1] && s2[j - 1] === s3[i + j - 1]); 
            }
        }
    }
    return dp[s1.length][s2.length];
};

var isInterleave = function(s1, s2, s3) {  // Based on the state transform equation, dp[i][j] relys on dp[i - 1][j] and dp[i][j - 1], so we can compress the space usage to 1D array
    if (s1.length + s2.length !== s3.length) {
        return false;   
    }
    const dp = [];
    for (let i = 0; i <= s1.length; i++) {
        for(let j = 0; j <= s2.length; j++) {
            if (i === 0 && j === 0) {
                dp[j] = true;   
            } else if (i === 0) {
                dp[j] = dp[j - 1] && s2[j - 1] === s3[i + j - 1];   
            } else if (j === 0) {
                dp[j] = dp[j] && s1[i - 1] === s3[i + j - 1];   
            } else {
                dp[j] = (dp[j - 1] && s2[j - 1] === s3[i + j - 1]) ||  (dp[j] && s1[i - 1] === s3[i + j - 1])l
            }
        }
    }
    return dp[s2.length];
    
}
