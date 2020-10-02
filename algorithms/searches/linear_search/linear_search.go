package linearsearch

// LinearSearch is a Searching function with complexity O(N)
func LinearSearch(arr []int, key int) int {
	// This "for" loop traverses the entire array and check
	// each of them if match
	for i := 0; i < len(arr); i++ {
		// The block that does the matching
		if arr[i] == key {
			return i
		}
	}
	// If the respective key not found
	// This function will return -1
	return -1
}
