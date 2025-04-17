package easy

func addBinary(a string, b string) string {
	res := make([]byte, 0)
	aP, bP := len(a)-1, len(b)-1

	var curNum, remainder int
	for aP >= 0 && bP >= 0 {
		aNum, bNum := int(a[aP]-'0'), int(b[bP]-'0')
		curNum, remainder = (aNum+bNum+remainder)%2, (aNum+bNum+remainder)/2

		res = append(res, byte(curNum)+'0')
		aP--
		bP--
	}

	if aP < 0 {
		a, aP = b, bP
	}

	for aP >= 0 {
		aNum := int(a[aP] - '0')
		curNum, remainder = (aNum+remainder)%2, (aNum+remainder)/2

		res = append(res, byte(curNum)+'0')
		aP--
	}

	if remainder != 0 {
		res = append(res, byte(remainder)+'0')
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return string(res)

}
