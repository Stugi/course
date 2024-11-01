package str

// Concat объединяет строки.
func Concat(ss ...string) string {
	var res string
	for _, s := range ss {
		res += s
	}
	return res
}

// Invert разворачивает строку.
func Invert(s string) string {
	// b := []byte(s) // только для unicode
	b := []rune(s)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	return string(b)
}
