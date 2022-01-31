package main

import "fmt"

func calcRemainder(numerator,denominator int) (int,int,error) {
	if denominator == 0 {
		return 0,0,fmt.Errorf("Division by zero") // uses fmt.Errorf()
	}

	return numerator/denominator, numerator%denominator, nil
}