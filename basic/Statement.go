package main

import (
	"fmt"
	"time"
)

func main() {

	var i int = 3

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}


	switch time.Now().Weekday() {
		case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
			fmt.Println("it's a weekday",time.Now().Weekday())
		case time.Saturday, time.Sunday:
			fmt.Println("it's a weekend",time.Now().Weekday())
	}

}
