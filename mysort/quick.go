package mysort

func Quicksort(s []int) []int {
	if len(s) == 1 || len(s) == 0 { //list size 0 or 1 is no sort
		return s
	} else {
		pivot := s[0] //make a pivot first in the list
		place := 0

		for j := 0; j < len(s)-1; j++ {
			if s[j+1] < pivot { // if it is smaller than the pivot
				s[j+1], s[place+1] = s[place+1], s[j+1]
				place++
			}
		}
		s[0], s[place] = s[place], s[0]

		first := Quicksort(s[:place])
		second := Quicksort(s[place+1:])
		first = append(first, s[place])

		first = append(first, second...)
		return first
	}

}
