package main

import "fmt"

func main() {

	// if we known the value for map we can use make.

	m := make(map[string]int)

	// dynamic map
	// mas := map[string]int {

	// }

	m["amazon"] = 1
	m["google"] = 2
	m["facebook"] = 3
	m["microsoft"] = 4

	fmt.Println(m)
	fmt.Println("map length = ",len(m))
	delete(m, "amazon")
	fmt.Println(m)

	val,ok := m["amazon"]
	fmt.Println(val,ok)

	employee := make(map[string]int)
	employee["john"] = 1
	employee["jane"] = 2
	employee["joe"] = 3
	fmt.Println(employee)

	val,ok = employee["jane"]

	 if ok == true {
		fmt.Println("jane is present in the map")
	 } else {
		fmt.Println("jane is not present in the map")
	 }
	




}

	





