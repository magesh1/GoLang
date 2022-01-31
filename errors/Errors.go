package main

import (
	// "errors"
	"fmt"
	"os"
)


func main() {
	num := 3
	denom := 0
	rem, mod, err := calcRemainder(num,denom)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(rem, mod)
	
	n, err := eveErr(2)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(n)
	

}


