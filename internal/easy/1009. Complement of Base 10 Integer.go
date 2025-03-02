package easy

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/complement-of-base-10-integer/

func bitwiseComplement(n int) int {
	const (
		one, zero byte = '1', '0'
	)

	var (
		bits = []byte(strconv.FormatInt(int64(n), 2))
		res  int64
	)

	for i := 0; i < len(bits); i++ {
		switch bits[i] {
		case one:
			bits[i] = zero

		case zero:
			bits[i] = one
		}
	}

	res, _ = strconv.ParseInt(string(bits), 2, 64)

	fmt.Println(res)

	return int(res)

}
