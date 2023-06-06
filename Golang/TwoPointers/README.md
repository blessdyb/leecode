## [Two pointers](https://beendless.com/2021/09/07/Two-Pointers/)

The Two Pointers algorithm is a technique commonly used to solve problems that involve searching, traversing, or manipulating arrays or linked lists. It involves using two pointers that move through the data structure in a specific manner to solve the problem efficiently.

The basic idea behind the Two Pointers algorithm is to use two pointers, often referred to as the "left" and "right" pointers, which are initially positioned at different indices of the array or linked list. These pointers can move towards each other or in different directions depending on the problem requirements. By manipulating the pointers based on certain conditions, you can find a solution or optimize the algorithm's time complexity.

Here's a step-by-step approach to using the Two Pointers algorithm to solve problems:

Initialize the pointers: Create two pointers, often referred to as "left" and "right," and set their initial positions. The starting positions depend on the problem's requirements.

Define the problem-specific conditions: Identify the conditions that determine how the pointers should move or behave. These conditions are problem-dependent and may involve comparisons with certain values or ranges.

Implement the algorithm: While the pointers have not crossed each other or until they reach the end of the data structure, repeatedly perform the following steps:

a. Check the condition(s): Evaluate the conditions specified for the problem using the values at the current positions of the pointers.

b. Adjust the pointers: Based on the condition(s), move the pointers to new positions. This step can involve moving both pointers closer, moving only one pointer, or adjusting their positions independently.

c. Perform necessary operations: Perform any necessary operations or computations based on the problem requirements. This step can involve updating variables, manipulating the data structure, or storing results.

Continue or terminate: Repeat steps 3a to 3c until the pointers meet, cross each other, or reach the end of the data structure, depending on the problem requirements.

Return the solution or result: After the algorithm terminates, return the desired solution or result based on the problem's requirements.

The key to using the Two Pointers algorithm effectively lies in understanding the problem and determining how to adjust the pointers based on the specific conditions. By manipulating the pointers strategically, you can often achieve better time complexity than alternative approaches.

```
func twoSum(nums []int, target int) bool {
    i, j := 0, len(nums) - 1
    for i <= j {
        sum := nums[i] + nums[j]
        if sum == target {
            return true
        } else if sum < target {
            i++
        } else {
            j--
        }
    }
    return false
}
```

Using two pointers with different speed (slower pointer and faster pointer), it can help us to resolve many array / link-list related problems. Especially for in-place updating problems which we have to make sure the space compexicity is O(1)

![image info](https://camo.githubusercontent.com/8c49666f4d541dd22e7e221964cc0ac7997d4bc98272d54ea6569c51bc2640d6/68747470733a2f2f636f64652d7468696e6b696e672e63646e2e626365626f732e636f6d2f676966732f32372e2545372541372542422545392539392541342545352538352538332545372542342541302d2545352538462538432545362538432538372545392539322538382545362542332539352e676966)


## Sliding Window

```
Note: even there are usually two nested loops, the time-complexity is not O(n^2). Since all elements will be visited by the sliding window at most twice (entering or exiting), so the time complexity is O(n) 
```

The sliding window algorithm is a technique used for efficiently solving problems that involve finding a substring, subarray, or subsequence that meets certain criteria within a larger string, array, or sequence. It is especially useful when the problem requires examining multiple contiguous elements at a time.

The basic idea behind the sliding window algorithm is to maintain a window (a subarray or substring) within the given data structure and slide it along the data structure while updating the window's contents based on specific conditions. By doing so, unnecessary computations can be avoided, leading to an optimized solution.

Here are the general steps involved in the sliding window algorithm:

Initialize the window boundaries: Set two pointers, left and right, to define the boundaries of the initial window. The window typically starts at the beginning of the data structure.

Expand the window: Move the right pointer to expand the window by including more elements in the current window. This step continues until the desired condition is met (e.g., a valid substring is found, or a target sum is achieved).

Shrink the window: If the current window satisfies the required condition, attempt to minimize the window size by moving the left pointer. This step continues until the condition is no longer satisfied. The objective is to find the minimum-sized window that still meets the problem criteria.

Update the solution: During the sliding process, track the best solution found so far based on the problem requirements. This could involve updating the minimum/maximum value, counting occurrences, or storing the substring or subarray itself.

Repeat steps 2-4: Continue expanding and shrinking the window, keeping track of the best solution, until the right pointer reaches the end of the data structure.

By using this sliding window approach, the algorithm reduces the overall time complexity by avoiding unnecessary computations and minimizing the number of elements being considered at each step.

![image info](https://camo.githubusercontent.com/dd84aee84237ebb78cf7ffde58803dc03350a4071d0981b8add65d9c59199ac4/68747470733a2f2f636f64652d7468696e6b696e672e63646e2e626365626f732e636f6d2f676966732f3230392e2545392539352542462545352542412541362545362539432538302545352542302538462545372539412538342545352541442539302545362539352542302545372542422538342e676966)