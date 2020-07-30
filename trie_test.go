package trie

import "testing"

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abc")
	trie.Insert("helloabc")
	trie.Insert("hellokkc")
	t.Log("init:", trie)

	t.Run("Test Exists", func(t *testing.T) {
		tests := map[string]bool{
			"abc":      true,
			"helloabc": true,
			"helloab":  false,
			"hellokkc": true,
		}
		for s, b := range tests {
			if trie.Exists(s) != b {
				t.Error(s)
			}
		}
	})
}
