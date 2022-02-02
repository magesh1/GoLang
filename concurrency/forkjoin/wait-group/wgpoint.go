package waitgroup

import (
	"fmt"
	"time"
)

func Work() {
	time.Sleep(500*time.Millisecond)
	fmt.Println("fork with join point...")
}