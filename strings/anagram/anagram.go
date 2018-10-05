package anagram

import "sort"

func IsAnagram(s1, s2 string) bool {
	return sortString(s1) == sortString(s2)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] > r[j]
	})

	return string(r)
}
