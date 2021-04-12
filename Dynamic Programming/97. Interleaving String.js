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

