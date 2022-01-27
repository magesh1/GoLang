package main

import "fmt"

func main() {
	 a := 21
	 if a > 10 {
		 a := 22 // shadowing
		 fmt.Println(a)
	 }
	 fmt.Println(a)
}
