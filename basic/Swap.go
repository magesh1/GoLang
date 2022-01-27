package main

import "fmt"

func main() {
	
	a := 12
	b := 13
	fmt.Println("before swaping")
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println("after swaping")
	fmt.Println(a,b)
	

}

func swap(a,b *int) {
	temp := *a
	*a = *b
	*b = temp

	
}
