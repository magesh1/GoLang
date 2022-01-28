package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 4, 3, 2, 1}
	n := 4
	shuffle(arr, n)
}

func shuffle(arr []int, n int) {
	res := make([]int, len(arr))
	store1 := make([]int, n)
	store2 := make([]int, n)

	for i := 0; i < n; i++ {
		store1[i] = arr[i]
	}

	k := 0

	for i := n; i < len(arr); i++ {
		store2[k] = arr[i]
		k += 1
	}

	k = 0

	for i := 0; i < len(store1); i++ {
		res[k] = store1[i]
		k += 2
	}

	k = 1

	for i := 0; i < len(store2); i++ {
		res[k] = store2[i]
		k += 2
	}

	fmt.Println(res)

}
