package easy

import (
	"strconv"
)

func freqAlphabets(s string) string {
	const (
		pad = 1

		sharp byte = '#'

		lowerStart byte = 'a' - pad
	)

	var (
		result = make([]byte, 0, len(s))

		reverse func(strBytes []byte) string = func(strBytes []byte) string {
			var (
				l = len(strBytes) / 2
			)

			for i := 0; i < l; i++ {
				strBytes[i], strBytes[len(strBytes)-i-pad] = strBytes[len(strBytes)-i-pad], strBytes[i]
			}

			return string(strBytes)
		}

		numBytes = make([]byte, 2)
	)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == sharp {
			i--

			for j := 0; j < 2; j++ {
				numBytes[j] = s[i]
				i--
			}

			doubleCharNum, _ := strconv.Atoi(reverse(numBytes))

			result = append(result, byte(doubleCharNum)+lowerStart)

			i++
			continue
		}

		singleCharNum, _ := strconv.Atoi(string(rune(s[i])))
		result = append(result, lowerStart+byte(singleCharNum))
	}

	return reverse(result)
}
