package strings

func RemoveColorCodes(s string) string {
	res := make([]byte, 0, len(s))

	i := 0
	for i < len(s) {
		switch {
		case s[i] > 127:
			break
		default:
			res = append(res, s[i])
		}
		i++
	}

	return string(res)
}
