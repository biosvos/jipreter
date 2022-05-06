package trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie(t *testing.T) {
	aa := []struct {
		s   string
		uid int
	}{
		{
			s:   "A",
			uid: 1,
		},
		{
			s:   "AB",
			uid: 2,
		},
		{
			s:   "ABC",
			uid: 3,
		},
	}
	for _, tt := range []struct {
		indices []int
	}{
		{indices: []int{0, 1, 2}},
		{indices: []int{0, 2, 1}},
		{indices: []int{1, 0, 2}},
		{indices: []int{1, 2, 0}},
		{indices: []int{2, 0, 1}},
		{indices: []int{2, 1, 0}},
	} {
		trie := NewTrie("", 0)

		for _, index := range tt.indices {
			_ = trie.Register(aa[index].s, aa[index].uid)
		}

		assert.Equal(t, 1, len(trie.children))
		assert.Equal(t, 1, trie.children[0].uid)
		assert.Equal(t, "A", trie.children[0].sub)
		assert.Equal(t, 1, len(trie.children[0].children))
		assert.Equal(t, 2, trie.children[0].children[0].uid)
		assert.Equal(t, "B", trie.children[0].children[0].sub)
		assert.Equal(t, 1, len(trie.children[0].children[0].children))
		assert.Equal(t, 3, trie.children[0].children[0].children[0].uid)
		assert.Equal(t, "C", trie.children[0].children[0].children[0].sub)
	}
}
