package trees

import (
	"testing"
)

func TestInsertAndSearchWord(t *testing.T) {
	trie := NewTrie()
	trie.insert("hello")
	if !trie.Search("hello") {
		t.Errorf("Expected to find 'hello' in the trie")
	}
}

func TestSearchNonExistentWord(t *testing.T) {
	trie := NewTrie()
	trie.insert("hello")
	if trie.Search("world") {
		t.Errorf("Did not expect to find 'world' in the trie")
	}
}

func TestInsertAndSearchPrefix(t *testing.T) {
	trie := NewTrie()
	trie.insert("hello")
	if !trie.StartsWith("hell") {
		t.Errorf("Expected to find prefix 'hell' in the trie")
	}
}

func TestDeleteWord(t *testing.T) {
	trie := NewTrie()
	trie.insert("hello")
	trie.Delete("hello")
	if trie.Search("hello") {
		t.Errorf("Did not expect to find 'hello' in the trie after deletion")
	}
}

func TestDeleteNonExistentWord(t *testing.T) {
	trie := NewTrie()
	trie.insert("hello")
	if trie.Delete("world") {
		t.Errorf("Did not expect to delete 'world' from the trie")
	}
}

func TestDeletePrefix(t *testing.T) {
	trie := NewTrie()
	trie.insert("hello")
	trie.Delete("hell")
	if !trie.Search("hello") {
		t.Errorf("Expected to find 'hello' in the trie after deleting prefix 'hell'")
	}
}
