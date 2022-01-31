package main

import "errors"


func eveErr(num int) (int,error) {
	if num % 2 == 0 {
		return 0,errors.New("Even number") // this one uses errors.New()
	}
	return num,nil
}