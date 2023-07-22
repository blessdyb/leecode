package main

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	pre, in := []string{}, []string{}
	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node != nil {
			pre = append(pre, strconv.Itoa(node.Val))
			preOrder(node.Left)
			preOrder(node.Right)
		}
	}
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node != nil {
			inOrder(node.Left)
			in = append(in, strconv.Itoa(node.Val))
			inOrder(node.Right)
		}
	}
	preOrder(root)
	inOrder(root)
	return strings.Join(pre, ",") + "#" + strings.Join(in, ",")
}

func (this *Codec) deserialize(data string) *TreeNode {
	if data == "#" {
		return nil
	}
	sequences := strings.SplitN(data, "#", 2)
	preOrder, inOrder := strings.Split(sequences[0], ","), strings.Split(sequences[1], ",")
	indexOf := func(slice []string, needle string) int {
		var index int
		for i, val := range slice {
			if val == needle {
				index = i
				break
			}
		}
		return index
	}

	var helper func(pre, in []string) *TreeNode
	helper = func(pre, in []string) *TreeNode {
		if len(pre) > 0 {
			root := pre[0]
			indexInInOrder := indexOf(in, root)
			leftInOrder := in[:indexInInOrder]
			leftPreOrder := pre[1 : len(leftInOrder)+1]
			rightInOrder := in[indexInInOrder+1:]
			rightPreOrder := pre[len(leftPreOrder)+1:]
			rootValue, _ := strconv.Atoi(root)
			return &TreeNode{
				Val:   rootValue,
				Left:  helper(leftPreOrder, leftInOrder),
				Right: helper(rightPreOrder, rightInOrder),
			}
		}
		return nil
	}
	return helper(preOrder, inOrder)
}
