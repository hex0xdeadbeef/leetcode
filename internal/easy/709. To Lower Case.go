package easy

func toLowerCase(s string) string {
	const diff byte = 'a' - 'A'

	sBytes := []byte(s)

	for i, b := range sBytes {
		if b >= 'a' && b <= 'z' {
			continue
		}

		if b >= 'A' && b <= 'Z' {
			sBytes[i] += diff
			continue
		}

	}

	return string(sBytes)
}
