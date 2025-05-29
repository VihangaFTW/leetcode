package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEndUnOptimized(head *ListNode, n int) *ListNode {

	//? space O(N) where N is the length of the list
	seenNodes := []*ListNode{}
	currentNode := head

	//? time O(N) where N is the length of the list
	for currentNode != nil {
		seenNodes = append(seenNodes, currentNode)
		currentNode = currentNode.Next
	}

	removeIndex := len(seenNodes) - n

	// edge case where we remove the only element in the list so return empty list
	if len(seenNodes) == 1 && removeIndex == 0 {
		return nil
	}

	//? remove the first element in linked list
	if removeIndex == 0 {
		return seenNodes[1]
	}

	//? link the previous node to the target node's next link; essentially breaking all references to the target node
	removeNode := seenNodes[removeIndex]

	seenNodes[removeIndex-1].Next = removeNode.Next

	return seenNodes[0]
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	
	//? dummy node so that when the right pointer is at nil, left pointer points to the node before the target node. Also handles the edge case where we remove the only element in the list
	dummy := ListNode{Val: 0, Next: head}

	left, right := &dummy, head

	//? keep the difference of n so that when right reaches the end of the list, left points to the node
	//? before the target node
	for range n {
		right = right.Next
	}


	//? move both pointers until  right reaches the end
	for right != nil {
		right = right.Next
		left = left.Next
	}

	//? remove the target node
	left.Next = left.Next.Next

	return dummy.Next

}

func TraverseList(head *ListNode) []int {
	nums := []int{}

	current := head

	for current != nil {
		nums = append(nums, current.Val)
		current = current.Next
	}

	return nums
}
