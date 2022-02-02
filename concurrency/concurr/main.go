package conctwo

import (
	"fmt"
	"time"
	"sync"
)



func Works(i int) {
	time.Sleep(100*time.Millisecond)
	fmt.Println("task", i, "completed")
}

func WaitGroup(wg *sync.WaitGroup,i int) {
	defer wg.Done()
	time.Sleep(100*time.Millisecond)
	fmt.Println("task", i, "completed")
}

func Deadlock() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}