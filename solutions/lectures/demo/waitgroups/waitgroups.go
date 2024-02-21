package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	sum := 0
	for i := 0; i < 30; i++ {
		wg.Add(1)
		x := i
		go func() {
			time.Sleep(time.Duration(10*x) * time.Millisecond)
			sum += x * x * x
			wg.Done()
			fmt.Printf("Result for %d calculated\n", x)
		}()
	}

	wg.Wait()
	fmt.Println("jobs complete")
	fmt.Println("sum =", sum)
}
