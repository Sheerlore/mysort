package mysort

func Insertionsort(s []int) []int {
	for i := 1; i < len(s); i++ {
		j := i - 1
		temp := s[i]
		for j > -1 && s[j] > temp {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = temp
	}
	return s
}
