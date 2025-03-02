package easy

// https://leetcode.com/problems/shortest-distance-to-target-string-in-a-circular-array/description/

func closetTarget(words []string, target string, startIndex int) int {
	var (
		isFound bool
	)

	for _, w := range words {
		if w == target {
			isFound = true
			break
		}
	}

	if !isFound {
		return -1
	}

	if words[startIndex] == target {
		return 0
	}

	const (
		pad = 1
	)

	var (
		minStepsCnt = len(words) + pad
	)

	switch curStepsCnt, ok := moveLeft(words, target, startIndex); {
	case ok:
		if curStepsCnt < minStepsCnt {
			minStepsCnt = curStepsCnt
		}
	default:
		additionalStepsCnt, _ := moveLeft(words, target, len(words)-1)
		if overallStepsCnt := curStepsCnt + additionalStepsCnt; overallStepsCnt < minStepsCnt {
			minStepsCnt = overallStepsCnt
		}
	}

	switch curStepsCnt, ok := moveRight(words, target, startIndex); {
	case ok:
		if curStepsCnt < minStepsCnt {
			minStepsCnt = curStepsCnt
		}
	default:
		additionalStepsCnt, _ := moveRight(words, target, 0)
		if overallStepsCnt := curStepsCnt + additionalStepsCnt; overallStepsCnt < minStepsCnt {
			minStepsCnt = overallStepsCnt
		}
	}

	return minStepsCnt
}

func moveLeft(words []string, target string, startIndex int) (int, bool) {
	var (
		stepsCnt int
		isFound  bool
	)

	for startIndex >= 0 {
		if words[startIndex] == target {
			isFound = true
			break
		}

		startIndex--
		stepsCnt++

	}

	return stepsCnt, isFound
}

func moveRight(words []string, target string, startIndex int) (int, bool) {
	var (
		stepsCnt int
		isFound  bool
	)

	for startIndex < len(words) {
		if words[startIndex] == target {
			isFound = true
			break
		}

		startIndex++
		stepsCnt++

	}

	return stepsCnt, isFound
}
