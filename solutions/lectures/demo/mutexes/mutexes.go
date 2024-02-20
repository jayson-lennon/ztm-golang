package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

type Hits struct {
	count int
	sync.Mutex
}

func serve(wg *sync.WaitGroup, hits *Hits, iteration int) {
	wait()
	hits.Lock()
	defer hits.Unlock()
	defer wg.Done()
	hits.count += 1
	fmt.Println("served iteration", iteration)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	hitCounter := Hits{}
	for i := 0; i < 20; i++ {
		iteration := i
		wg.Add(1)
		go serve(&wg, &hitCounter, iteration)
	}

	fmt.Printf("Waiting for goroutines...\n\n")
	wg.Wait()

	hitCounter.Lock()
	totalHits := hitCounter.count
	hitCounter.Unlock()

	fmt.Printf("\ntotal hits = %d\n", totalHits)
}
