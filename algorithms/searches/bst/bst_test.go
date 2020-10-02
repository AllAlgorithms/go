package main

import "testing"

// This is called a test table. It's a way to easily
// specify tests while avoiding boilerplate.
// See https://github.com/golang/go/wiki/TableDrivenTests
var tests = []struct {
	input  int
	output bool
}{
	{6, true},
	{16, false},
	{3, true},
}

func TestSearch(t *testing.T) {
	//     6
	//    /
	//   3
	tree := &Node{Key: 6, Left: &Node{Key: 3}}

	for i, test := range tests {
		if res := tree.Search(test.input); res != test.output {
			t.Errorf("%d: got %v, expected %v", i, res, test.output)
		}
	}

}

func BenchmarkSearch(b *testing.B) {
	tree := &Node{Key: 6}

	for i := 0; i < b.N; i++ {
		tree.Search(6)
	}
}
