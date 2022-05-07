package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func fpWork(wg *sync.WaitGroup, tid int) {
	t0 := time.Now()
	var sum float64
	sum = 0
	for i := 0; i < 123456; i = i + 1 {
		for j := 0; j < 123456; j = j + 1 {
			sum += float64(i) * float64(j) / 1000.0
		}
	}

	diff := time.Now().Sub(t0)
	fmt.Println("goroutine", tid, "finished. Float64, Value =", sum, "Total Time took:", diff)
	wg.Done()
}

func intWork(wg *sync.WaitGroup, tid int) {
	t0 := time.Now()
	var sum, i, j int64
	sum = 0
	for i = 0; i < 123456; i++ {
		for j = 0; j < 123456; j++ {
			sum += i * j / 1000
		}
	}

	diff := time.Now().Sub(t0)
	fmt.Println("goroutine", tid, "finished. Int64, Value =", sum, "Total Time took:", diff)
	wg.Done()
}

func main() {
	concurrency := (flag.Int("c", 2, "concurrency level"))
	flag.Parse()
	fmt.Println("hello, benchmarking FP + INT with", *concurrency, "goroutine(s)")
	var wg sync.WaitGroup
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go fpWork(&wg, i)
	}
	wg.Wait()

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go intWork(&wg, i)
	}
	wg.Wait()
}
