## Binary Search

To identify a binary search problem, the array in the given problem must be a asending & sorted. If we are asked to resolve a searching task in O(log(n)), in most cases, binary search is the one we need to choose from our toolbox.

When you start writing down a binary search solution, always keep the open/close interval in mind and keep it consistant through your whole code base. 
* If you want to use [0, len - 1] as the range, i <= j needs to be chosen as the condition
* For [0, len] range selection, i < j will be the comparasion you need to use.

```
func binarySearch([]int nums, int target) int {
    i, j := 0, len(nums) - 1
    for (i <= j) {
        h := i + (j - i) / 2
        if (nums[h] > target) {
            j = h - 1
        } else (nums[h] < target) {
            i = h + 1
        } else {
            return h
        }
    }
    return -1
}
```