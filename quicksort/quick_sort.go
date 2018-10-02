package quicksort

func Sort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	pivot := slice[len(slice)-1]
	i, j := 0, 0
	for count := len(slice); i < count; i++ {
		if slice[i] <= pivot {
			slice[j], slice[i] = slice[i], slice[j]
			j++
		}
	}

	left, right := Sort(slice[:j-1]), Sort(slice[j:])
	return append(append(left, pivot), right...)
}
