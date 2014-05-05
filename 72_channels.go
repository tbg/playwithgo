package main

import (
	"fmt"

	"code.google.com/p/go-tour/tree"
)

func DoWalk(t1, t2 *tree.Tree, ch1, ch2 chan int, done chan bool) {
	fmt.Println("Walking")
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	fmt.Println("Triggering control channel")
	done <- true
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	fmt.Println("Writing to channel")
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var ch1, ch2 = make(chan int), make(chan int)
	var done = make(chan bool)
	go DoWalk(t1, t2, ch1, ch2, done)
	<-done

	var b int
	var ok bool = true
	for a := range ch1 {
		fmt.Println("Read", a)
		if !ok {
			// b was done in the last iteration, but a has another item -> not equal
			fmt.Println("Tree 1 has more elements")
			return false
		}
		b, ok = <-ch2
		if a != b {
			fmt.Println("Different elements encountered")
			return false
		}
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
