package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

//helper function to handle the recursion of the Walk function
func walkHelper(t *tree.Tree, ch chan int) {
	// Handle left tree
	if t.Left != nil {
		walkHelper(t.Left, ch)
	}

	// Send the value to the channel
	ch <- t.Value

	// Handle the right tree
	if t.Right != nil {
		walkHelper(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// Using the helper for the recursion allows me to close
// the channel when it comes back from recursing
func Walk(t *tree.Tree, ch chan int) {
	walkHelper(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// Go through each elt in the first tree
	for x := range ch1 {
		y, ok := <-ch2

		// If the first tree is longer they !=
		if !ok || x != y {
			return false
		}
	}

	// If the second tree is longer they !=
	if _, ok := <-ch2; ok {
		return false
	} else {
		return true
	}
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

