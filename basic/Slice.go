package main

import (
	"fmt"
	"sort"
)

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("len = ", len(s))

	s = append(s, "d")
	fmt.Println("new set:", s)

	fmt.Println("new len = ", len(s))

	s = append(s, "e", "f")
	fmt.Println("new set:", s)

	fmt.Println("new len = ", len(s))

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("cpy:", c)
	c = append(c, "g")
	fmt.Println("new cpy:", c)

	split := c[:4]
	fmt.Println("split:", split)

	arr := []int{10, 21, 11, 3, 1, 7, 6}
	sort.Ints(arr)
	fmt.Println(arr)

}
