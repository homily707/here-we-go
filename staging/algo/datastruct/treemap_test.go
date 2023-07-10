package datastruct

import (
	"fmt"
	"strings"
	"testing"
)

func Test_TreeMap(t *testing.T) {
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"

	root := NewTrieNode()
	for _, word := range dictionary {
		root.Add(word)
	}
	ss := strings.Split(sentence, " ")
	builder := strings.Builder{}
	for _, v := range ss {
		fmt.Print(v)
		builder.WriteString(root.SearchPre(v)[0])
	}
	fmt.Println(builder.String())
}
