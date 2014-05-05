package main

import "fmt"

func Worker(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
}

func main() {
	var ch = make(chan int)
	go Worker(ch)
	var elem int
	var ok bool
	for {
		elem, ok = <-ch
		fmt.Println(elem, ok)
		if !ok {
			break
		}
	}
}
