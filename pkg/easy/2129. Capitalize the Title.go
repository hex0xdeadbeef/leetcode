package easy

import (
	"bytes"
)

// https://leetcode.com/problems/capitalize-the-title/

func сapitalizeTitle(title string) string {
	const (
		pad byte = 1 << 5
	)
	var (
		strBytes = bytes.Split([]byte(title), []byte{' '})
	)

	for i := 0; i < len(strBytes); i++ {

		strBytes[i] = bytes.ToLower(strBytes[i])

		if len(strBytes[i]) > 2 {
			strBytes[i][0] -= pad
		}

	}

	return string(bytes.Join(strBytes, []byte{' '}))
}
