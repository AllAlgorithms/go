package kfrequentwords

import (
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func Top(s string, k int) []string {
	m := map[string]int{}
	for _, w := range strings.Fields(s) {
		m[w] += 1
	}

	wcs := []wordCount{}
	for w, c := range m {
		wc := wordCount{w, c}
		wcs = append(wcs, wc)
	}

	sort.Slice(wcs, func(i, j int) bool {
		return wcs[i].count > wcs[j].count
	})

	if k > len(wcs) {
		k = len(wcs)
	}

	topWords := []string{}
	for _, wc := range wcs[:k] {
		topWords = append(topWords, wc.word)
	}

	return topWords
}
