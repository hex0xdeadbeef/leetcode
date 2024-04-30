package easy

// https://leetcode.com/problems/roman-to-integer/description/

func RomanToInt(s string) int {
	const (
		indPad = 1
	)

	var (
		romanToInt = map[string]int{
			"I": 1, "IV": 4, "V": 5, "IX": 9,
			"X": 10, "XL": 40, "L": 50, "XC": 90,
			"C": 100, "CD": 400, "D": 500,
			"CM": 900, "M": 1_000}

		result int
	)

	for i := 0; i < len(s); i++ {
		switch i+indPad < len(s) {
		case false:
			result += romanToInt[string(s[i])]
		case true:
			if val, ok := romanToInt[s[i:i+indPad+indPad]]; ok {
				result += val
				i++
				continue
			}
			result += romanToInt[string(s[i])]
		}
	}

	return result

}
