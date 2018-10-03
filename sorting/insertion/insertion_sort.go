package insertion

/*
 * Insertion sort - https://en.wikipedia.org/wiki/Insertion_sort
 */

func Sort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	var i, j int
	for i = 1; i < len(arr); i++ {
		for j = 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
