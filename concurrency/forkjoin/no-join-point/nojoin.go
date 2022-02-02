package nojoins

import (
	"fmt"
	"time"
)


func Work() {
	time.Sleep(700*time.Millisecond)
	fmt.Println("fork without join point...")
}