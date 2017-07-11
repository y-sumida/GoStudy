package main

import "fmt"

func fibonacci() func() int {
	n := 0
	m := 1
	l := 0

	return func() int {
		result := 0

		if n == 0 {
			result = 0
		} else {
			result = l
		}

		l = n + m
		n = m
		m = l

		return result
	}


}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

