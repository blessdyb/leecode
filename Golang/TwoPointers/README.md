## Two pointers

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