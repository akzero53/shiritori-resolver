package utils

func GetStringAt(s string, i int) string {
	rs := []rune(s)
	return string(rs[i])
}
