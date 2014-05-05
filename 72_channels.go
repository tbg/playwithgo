package main

import (
	"fmt"
	"time"

	"code.google.com/p/go-tour/tree"
)

func DoWalk(t1, t2 *tree.Tree, ch1, ch2 chan int) {
	fmt.Println("Walking")
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	fmt.Println("Writing to channel:", t.Value)
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var ch1, ch2 = make(chan int), make(chan int)
	go DoWalk(t1, t2, ch1, ch2)

	var b int
	var ok bool
	for a := range ch1 {
		fmt.Println("Read from Tree 1:", a)
		b, ok = <-ch2
		if !ok {
			// b is done, but a has another item -> not equal
			fmt.Println("Tree 1 has more elements")
			return false
		}
		fmt.Println("Read from Tree 2:", b)
		if a != b {
			fmt.Println("Different elements encountered")
			return false
		}
		fmt.Println("Elements match")
		time.Sleep(1e9)
	}
	select {
	case _, ok := <-ch2:
		if ok {
			fmt.Println("Tree 2 has more elements")
			return false
		}
	}
	return true

}

func main() {
	if Same(tree.New(2), tree.New(2)) {
		fmt.Println("Good")
	} else {
		fmt.Println("Bad")
	}
}
