package iteration

func Repeat(c string, count int) string {
	result := ""

	if count > 0 {
		for i := 0; i < count; i++ {
			result += c
		}
	}

	return result
}
