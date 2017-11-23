package main

import (
	"fmt"
)

//fibonacci is a function that returns
//a function that returns an int.
func fibonacci() func() int {
	call := 0
	prev := 0
	last := 1
	return func() int {
		switch call {
		case 0:
			call += 1
			return 0
		case 1:
			call += 1
			return 1
		}
		prev, last = last, prev+last
		return last
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
