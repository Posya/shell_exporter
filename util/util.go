package util

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func FilterSlice(s []string, f func(string) string) []string {
	var result []string
	for _, str := range s {
		if r := f(str); r != "" {
			result = append(result, r)
		}
	}
	return result
}
