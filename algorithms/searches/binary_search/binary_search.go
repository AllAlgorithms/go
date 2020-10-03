package binarysearch

// BinarySearch is a Searching function for Ascending Sorted Data.
func BinarySearch(arr []int, left int, right int, key int) int {
	if right >= left {
		mid := left + (right-left) / 2

		// If the element is present at the middle itself
		if arr[mid] == key {
			return mid
		}

		// If element is smaller than mid, then
		// it can only be present in left subarray
		if arr[mid] > key {
			return BinarySearch(arr, left, mid-1, key)
		}

		// Else the element can only be present
		// in right subarray
		return BinarySearch(arr, mid+1, right, key)
	}

	// We reach here when element is not
	// present in array
	return -1
}
