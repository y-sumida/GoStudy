package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var f func(*tree.Tree)
	f = func(t *tree.Tree) {
		if t.Left != nil {
			f(t.Left)
		}
		ch <- t.Value

		if t.Right != nil {
			f(t.Right)
		}
	}
	f(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go Walk(t1, ch1)
    go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		n1 := <- ch1
		n2 := <- ch2

		if n1 != n2 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
