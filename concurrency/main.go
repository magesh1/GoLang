package main

import (
	// "fmt"
	// "runtime"
	// "sync"
	// "time"

	// "github.com/magesh1/GoLang/concurrency/sync"
	// "github.com/magesh1/GoLang/concurrency/forkjoin/no-join-point"
	"github.com/magesh1/GoLang/concurrency/concurr"
	// "github.com/magesh1/GoLang/concurrency/forkjoin/wait-group"
)

// go keyword is used for concurrent execution of multiple go routines and became async

func main() {
	// now := time.Now()

	/**
	// sync and async tasks
	go syncs.Task1()
	go syncs.Task2()
	go syncs.Task3()
	go syncs.Task4()
	time.Sleep(time.Second)
	fmt.Println("Starting at:", time.Since(now))
	**/
	
	/**
	// no join point
	go nojoins.Work() // fork point
	time.Sleep(100*time.Millisecond)
	fmt.Println("done waiting,main exiting")
	**/
	
	/** 
	// wait group
	now := time.Now()
	var wg sync.WaitGroup // how many tasks to wait for (waitgroup) this keyword is used
	wg.Add(1) // we have only one task to wait for so the value is 1
	go func() { // anonymous function 
		defer wg.Done() // done waiting
		waitgroup.Work()
	} ()
	wg.Wait() // fork point to wait for all tasks to complete	
	fmt.Println("Starting at:", time.Since(now))
	fmt.Println("done waiting,main exiting")
	**/

	/**
	//concurr package
	fmt.Println("number of cores: ",runtime.NumCPU())
	
	for i := 0;i < 10;i++ {
		go conctwo.Works(i)
	}

	time.Sleep(time.Second)
	fmt.Println("done waiting,main exiting")
	**/
/**
	//using waitgroup
	fmt.Println("number of cores: ",runtime.NumCPU())
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0;i < 10;i++ {
		go conctwo.WaitGroup(&wg,i)
	}
	wg.Wait()
	fmt.Println("Starting at:", time.Since(now))
	fmt.Println("done waiting,main exiting")
**/

	//deadlock
	conctwo.Deadlock()


}