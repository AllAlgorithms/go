package bucketsort

/*
	Author: poifull10
*/

func Sort(values []int) []int {

	/* First, obtain max values of the given array. */
	max := 0
	for _, e := range values {
		if max < e {
			max = e
		}
	}

	/* Preparation of buckets. */
	buckets := make([]bool, max+1)

	for _, e := range values {
		buckets[e] = true
	}

	ret := make([]int, len(values))
	t := 0

	for i, e := range buckets {
		if e {
			ret[t] = i
			t++
		}
	}

	return ret
}
