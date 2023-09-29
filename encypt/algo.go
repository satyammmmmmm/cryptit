package encypt

func Nimbus(str string) string {
	encStr := ""
	for _, c := range str {
		asc := int(c)
		ch := string(asc + 3)
		encStr += ch
	}
	return encStr
}
