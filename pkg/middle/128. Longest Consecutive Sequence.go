package middle

// https://leetcode.com/problems/longest-consecutive-sequence/description/

func longestConsecutive(nums []int) int {
	var (
		occurencies = make(map[int]empty, len(nums))

		head, tail     int
		curCnt, maxCnt int
	)

	for _, v := range nums {
		occurencies[v] = empty{}
	}

	if len(occurencies) == 0 && len(occurencies) == 1 {
		return len(occurencies)
	}

	tail, head, curCnt = nums[0], nums[0], 1
	delete(occurencies, tail)

	for len(occurencies) != 0 {
		if _, ok := occurencies[tail+1]; ok {
			tail = tail + 1
			delete(occurencies, tail)
			curCnt++
		}

		if _, ok := occurencies[head-1]; ok {
			head = head - 1
			delete(occurencies, head)
			curCnt++
		}

		_, ok1 := occurencies[tail+1]
		_, ok2 := occurencies[head-1]

		if !ok1 && !ok2 {
			if curCnt > maxCnt {
				maxCnt = curCnt
			}

			for k, _ := range occurencies {
				tail, head, curCnt = k, k, 1
				delete(occurencies, k)
				break
			}

		}

	}

	return maxCnt

}
