package middle

// https://leetcode.com/problems/copy-list-with-random-pointer/description/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var (
		newList  = copyList(head)
		oldToNew = getOldToNewMap(head, newList)
	)

	for k, v := range oldToNew {
		if k.Random == nil {
			continue
		}

		randElemOfOld := k.Random
		v.Random = oldToNew[randElemOfOld]
	}

	return newList
}

func copyList(root *Node) *Node {

	if root == nil {
		return nil
	}

	var (
		oldCur = root

		newHead = &Node{Val: -1, Next: &Node{Val: -1}}
		newCur  = newHead.Next
	)

	for oldCur != nil {

		newCur.Val, newCur.Next = oldCur.Val, &Node{Val: -1}
		if oldCur.Next == nil {
			newCur.Next = nil
			break
		}

		oldCur = oldCur.Next
		newCur = newCur.Next
	}

	return newHead.Next
}

func getOldToNewMap(old, new *Node) map[*Node]*Node {
	var (
		m = make(map[*Node]*Node, 1<<9)
	)

	for old != nil {
		m[old] = new

		old, new = old.Next, new.Next
	}

	return m
}
