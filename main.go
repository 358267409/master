// TestChan project main.go
package main

import (
	"fmt"
)

const ngoroutine = 3

func func1(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	//	fmt.Println(right)
	//	fmt.Println(left)
	for i := 0; i < ngoroutine; i++ {
		left, right = right, make(chan int)
		//		fmt.Println(right)
		//		fmt.Println(left)
		go func1(left, right)
	}

	right <- 0
	fmt.Println(<-leftmost)

	var v1 *int
	fmt.Printf("%T,%v\n", v1, v1)

	var a, b, _, d, _ int = 1, 2, 3, 4, 5
	fmt.Printf("%v,%v,%v\n", a, b, d)

}
