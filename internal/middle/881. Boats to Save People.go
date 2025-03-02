package middle

// https://leetcode.com/problems/boats-to-save-people/description/?envType=daily-question&envId=2024-05-04

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {

	const (
		maxSeatsTakenCount = 2
		idxPad             = 1
	)

	var (
		l, r int = 0, len(people) - idxPad

		curLimit int = limit

		boatsCount, takenSeats int
	)

	sort.Ints(people)

	for l < r {
		if curLimit-people[r] >= 0 {
			curLimit -= people[r]

			takenSeats++
			r--
		}

		if curLimit-people[l] >= 0 {
			curLimit -= people[l]

			takenSeats++
			l++
		}

		if takenSeats == maxSeatsTakenCount || curLimit-people[r] < 0 && curLimit-people[l] < 0 || curLimit == 0 {
			boatsCount++

			takenSeats = 0
			curLimit = limit
		}
	}

	if l == r {
		boatsCount++
	}

	return boatsCount
}
