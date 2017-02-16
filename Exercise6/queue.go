package main

import (
	"fmt"
)

func main() {
	q := make([]int, 1)

	q = append(q, 1)

	fmt.Println(q)

	q = q[1:]

	fmt.Println(q)

	q = append(q, 2)

	fmt.Println(q)
}
