package middle

// https://leetcode.com/problems/copy-list-with-random-pointer/description/

type NodeWithRandom struct {
	Val    int
	Next   *NodeWithRandom
	Random *NodeWithRandom
}

func copyRandomList(head *NodeWithRandom) *NodeWithRandom {
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

func copyList(root *NodeWithRandom) *NodeWithRandom {

	if root == nil {
		return nil
	}

	var (
		oldCur = root

		newHead = &NodeWithRandom{Val: -1, Next: &NodeWithRandom{Val: -1}}
		newCur  = newHead.Next
	)

	for oldCur != nil {

		newCur.Val, newCur.Next = oldCur.Val, &NodeWithRandom{Val: -1}
		if oldCur.Next == nil {
			newCur.Next = nil
			break
		}

		oldCur = oldCur.Next
		newCur = newCur.Next
	}

	return newHead.Next
}

func getOldToNewMap(old, new *NodeWithRandom) map[*NodeWithRandom]*NodeWithRandom {
	var (
		m = make(map[*NodeWithRandom]*NodeWithRandom, 1<<9)
	)

	for old != nil {
		m[old] = new

		old, new = old.Next, new.Next
	}

	return m
}

func сopyRandomListRepeat(head *NodeWithRandom) *NodeWithRandom {
	var cur1 *NodeWithRandom
	cur2 := &NodeWithRandom{}
	newHead := cur2

	cur1 = head
	for ; cur1 != nil; cur1 = cur1.Next {
		copiedNode := &NodeWithRandom{Val: cur1.Val}
		cur2.Next, cur2 = copiedNode, copiedNode
	}
	newHead = newHead.Next

	cur1, cur2 = head, newHead
	m := make(map[*NodeWithRandom]*NodeWithRandom, 1<<8)
	for ; cur1 != nil; cur1, cur2 = cur1.Next, cur2.Next {
		m[cur1] = cur2
	}

	for k, v := range m {
		if k.Random == nil {
			continue
		}
		v.Random = m[k.Random]
	}

	return newHead
}
