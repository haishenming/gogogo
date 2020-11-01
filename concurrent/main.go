package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		fmt.Println("1")
		wg.Done()
	}()

	fmt.Println("2")
	wg.Wait()
}
