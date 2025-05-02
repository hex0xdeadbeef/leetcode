package middle

import . "leetcode/pkg/datastructures/linkedlist"

	func splitListToParts(head *ListNode, k int) []*ListNode {
		head = ReverseLinkedList(head)

		cur := head

		var l int
		for ; cur != nil; cur = cur.Next {
			l++
		}

		partL := l / k

		if partL == 0 {
			partL = 1
		}

		cur = head
		var (
			res      []*ListNode
			curCount int
		)

		for cur != nil {
			t, h := cur, cur
			for i := 0; t.Next != nil && i < partL-1; i++ {
				h = h.Next
			}

			curCount++

			if curCount == k && (float64(l)/float64(k)-float64(l/k)) > 0 {
				for h.Next != nil {
					h = h.Next
				}

				res = append(res, t)
				break
			}

			cur, h.Next = h.Next, nil
			res = append(res, t)

		}

		for i := range len(res) {
			res[i] = ReverseLinkedList(res[i])
		}

		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}

		for range k - l {
			res = append(res, nil)
		}

		return res
	}
