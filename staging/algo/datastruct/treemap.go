package datastruct

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
)

func redblack() {
	tree := redblacktree.NewWithIntComparator()
	tree.Put(1, "x") // 1->x
	tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	fmt.Println(tree.String())
	tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	fmt.Println(tree.String())
	tree.Put(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)
	fmt.Println(tree.String())
	tree.Put(2, "b") // 1->x, 2->b (in order)
	fmt.Println(tree.String())
	tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
	fmt.Println(tree.String())
	tree.Put(7, "c") // 1->a, 2->b, 3->c (in order)
	fmt.Println(tree.String())

	tree.Get(5)
}

func treeMap() {
}
