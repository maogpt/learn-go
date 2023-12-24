package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}

		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}

	walker(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1 {
		if v2, ok := <-ch2; !ok || v1 != v2 {
			return false
		}
	}

	_, ok := <-ch2
	return !ok
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)

	fmt.Println(Same(t1, t2))

	t3 := tree.New(1)
	t4 := tree.New(2)

	fmt.Println(Same(t3, t4))
}
