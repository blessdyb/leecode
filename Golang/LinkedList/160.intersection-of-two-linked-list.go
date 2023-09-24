package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var ret *ListNode
	p1 := headA
	for p1 != nil {
		p2 := headB
		for p2 != nil {
			if p1 == p2 {
				ret = p1
				return ret
			}
			p2 = p2.Next
		}
		p1 = p1.Next
	}
	return ret
}

func getIntersectionNodeHash(headA, headB *ListNode) *ListNode {
	hashmap := map[*ListNode]bool{}
	for headA != nil {
		hashmap[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if hashmap[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func getIntersectionNodeTwoPointer(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
