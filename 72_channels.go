package main

import (
	"fmt"

	"code.google.com/p/go-tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var ch1, ch2 = make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	var b int
	var ok bool
	for a := range ch1 {
		fmt.Println(a)
		b, ok = <-ch2
		if !ok || a != b {
			return false
		}

	}
	return true

}

func main() {
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("Good")
	}
}