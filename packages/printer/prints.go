package prints

import (
	// "fmt"
	"github.com/kyokomi/emoji/v2"
)

func PrintList(num int,flag bool) {
	if flag == true {
		// fmt.Printf("the given number %d is even :pizza: \n",num)
		emoji.Printf("the given number %d is even :tada: \n",num)
	} else {
		// fmt.Printf("the given number %d is odd\n",num)
		emoji.Printf("the given number %d is odd :unamused: \n",num)
	}
}