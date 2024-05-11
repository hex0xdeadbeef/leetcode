package easy

// https://leetcode.com/problems/check-if-n-and-its-double-exist/

func CheckIfExist(arr []int) bool {
	var (
		nums = make(map[int]empty, len(arr))

		zerosCount int
	)

	for _, v := range arr {
		if v == 0 {
			zerosCount++
		}
		nums[v] = empty{}
	}

	for k, _ := range nums {

		if k == 0 && zerosCount >= 2 {
			return true
		}

		if _, ok := nums[k/2]; k != 0 && ok && k%2 == 0 {
			return true
		}

		if _, ok := nums[k*2]; k != 0 && ok {
			return true
		}
	}

	return false

}
