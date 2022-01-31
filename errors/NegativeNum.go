package main

import "fmt"

func sqrt(f float64) (float64, error) {

	if f < 0 {
		return 0, fmt.Errorf("%v is negative", f) // %v represents the value of the variable
	}

	return f * 2, nil
}
