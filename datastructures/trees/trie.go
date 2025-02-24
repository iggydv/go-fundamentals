package trees

// Trie is a tree data structure that stores a dynamic set of strings.
type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// NewTrie creates a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

func (t *Trie) insert(word string) {
	current := t.root

	for _, char := range word {
		if _, ok := current.children[char]; !ok {
			current.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		current = current.children[char]
	}
	current.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.searchNode(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.searchNode(prefix) != nil
}

func (t *Trie) Delete(word string) bool {
	return t.deleteHelper(t.root, word, 0)
}

func (t *Trie) searchNode(str string) *TrieNode {
	current := t.root

	for _, char := range str {
		if _, exists := current.children[char]; !exists {
			return nil
		}
		current = current.children[char]
	}
	return current
}

func (t *Trie) deleteHelper(node *TrieNode, word string, index int) bool {
	if index == len(word) {
		if !node.isEnd {
			return false
		}
		node.isEnd = false
		return len(node.children) == 0
	}

	char := rune(word[index])
	if _, ok := node.children[char]; !ok {
		return false
	}

	child := node.children[char]
	shouldDelete := t.deleteHelper(child, word, index+1)

	if shouldDelete {
		delete(node.children, char)
		return len(node.children) == 0
	}
	return false
}
