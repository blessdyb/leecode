/**
 * @param {number} n
 * @return {number}
 */
var countVowelStrings = function(n) {  // Recursive
    return (function counter(previousIndex, leftDigits) {
        if (leftDigits > 0) {
            let count = 0;
            for(let i = 0; i < 5; i++) {
                 count += counter(i, leftDigits - 1);
            }
            return count;
        }
        return 1; // Base case, given prefix character, the total combination is 1
    })(0, n);
};
