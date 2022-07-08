package datastruct

type Trie interface {
	Search(word string) []string
	Add(word string)
}

type trieNode struct {
	next    []*trieNode
	visited bool
	isEnd   bool
	word    string
}

func NewTrieNode() *trieNode {
	root := &trieNode{
		make([]*trieNode, 26), true, false, "",
	}
	return root
}

func (t *trieNode) Add(word string) {
	for i := 0; i < len(word); i++ {
		index := getIndexFromLetter(word[i])
		if t.next == nil {
			t.next = make([]*trieNode, 26)
		}
		if t.next[index] == nil {
			t.next[index] = &trieNode{
				make([]*trieNode, 26), true, false, "",
			}
		}
		if i == len(word)-1 {
			t.next[index].isEnd = true
			t.next[index].word = word
		}
		t = t.next[index]
	}
}

func (t *trieNode) SearchPre(word string) []string {
	var ans []string
	for i := 0; i < len(word); i++ {
		index := getIndexFromLetter(word[i])
		if t.next == nil || t.next[index] == nil || !t.next[index].visited {
			ans = append(ans, word)
			return ans
		}
		if t.next[index].isEnd {
			ans = append(ans, t.next[index].word)
		}
		t = t.next[index]
	}
	ans = append(ans, word)
	return ans
}

func getIndexFromLetter(b byte) int {
	return (int)(b - 'a')
}
