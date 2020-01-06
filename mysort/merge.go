package mysort

func Mergesort(s []int) []int {
	var result []int
	if len(s) < 2 {
		return s
	}

	mid := int(len(s) / 2)
	r := Mergesort(s[:mid])
	l := Mergesort(s[mid:])
	i, j := 0, 0

	for i < len(r) && j < len(l) {
		if r[i] > l[j] {
			result = append(result, l[j])
			j++
		} else {
			result = append(result, r[i])
			i++
		}
	}

	result = append(result, r[i:]...)
	result = append(result, l[j:]...)

	return result
}
