package main

func pathInZigZagTree(label int) []int {
	type Node struct {
		Val    int
		Left   *Node
		Right  *Node
		Parent *Node
	}
	level := 1
	stack := []*Node{&Node{Val: 1}}
	n := 2
	labelNode := stack[0]
	for len(stack) > 0 {
		child := []*Node{}
		if level%2 == 1 {
			for i := len(stack) - 1; i >= 0; i-- {
				if n <= label {
					stack[i].Left = &Node{
						Val:    n,
						Parent: stack[i],
					}
					n++
					child = append(child, stack[i].Left)
					labelNode = stack[i].Left
				}
				if n <= label {
					stack[i].Right = &Node{
						Val:    n,
						Parent: stack[i],
					}
					n++
					child = append(child, stack[i].Right)
					labelNode = stack[i].Right
				}
			}
		} else {
			for i := 0; i < len(stack); i++ {
				if n <= label {
					stack[i].Left = &Node{
						Val:    n,
						Parent: stack[i],
					}
					n++
					child = append(child, stack[i].Left)
					labelNode = stack[i].Left
				}
				if n <= label {
					stack[i].Right = &Node{
						Val:    n,
						Parent: stack[i],
					}
					n++
					child = append(child, stack[i].Right)
					labelNode = stack[i].Right
				}
			}
		}
		if n > label {
			break
		}
		stack = child
	}
	ret := []int{}
	for labelNode != nil {
		ret = append([]int{labelNode.Val}, ret...)
		labelNode = labelNode.Parent
	}
	return ret
}
