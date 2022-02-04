package main

import (
	"fmt"
	"os"
	"path"
)

type Phone struct {
	Name string
	Phno int
}

var data = []Phone{}



func main() {
	ar := os.Args

	if len(ar) < 2 {
		exe := path.Base(ar[0])
		fmt.Printf("Usage: %s search|list <arguments>\n",exe)
		return
	}

	data = append(data, Phone{"jack",8976756454})
	data = append(data, Phone{"john",12345353432})
	data = append(data, Phone{"mary",12234543643})
	data = append(data, Phone{"jai",12234543233})

	switch ar[1] {
	case "search":
		if len(ar) != 3 {
			fmt.Println("usage: search name")
			return
		}

		result := search(ar[2])

		if result == nil {
			fmt.Println("username not found",ar[2])
			return
		}

		fmt.Println(*result)

	case "list":
		list()

	default:
		fmt.Println("invalid command")
	}

}

func search(name string) *Phone {
	for i, v := range data {
		if v.Name == name {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _,v := range data {
		fmt.Println(v)
	}
}