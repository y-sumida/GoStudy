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
// func Same(t1, t2 *tree.Tree) bool

func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)

	for i := range ch {
		fmt.Println(i)
	}
}
