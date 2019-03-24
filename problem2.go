package main

import (
	"sync"
	"log"
	"math/rand"
	"time"
)

func problem2() {

	log.Printf("problem2: started --------------------------------------------")

	jobs := make(chan int, 5)
	numbers := make(chan float32, 5)
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	var wg sync.WaitGroup
	wg.Add(5)

	for inx := 0; inx < 10; inx++ {

		go printRandom2(inx, &wg, jobs, numbers)

	}

	wg.Wait()
	close(numbers)

	limiter := time.Tick(1 * time.Second)

	nums := make([]float32, 0)
	for n := range numbers {
		<-limiter
		nums = append(nums, n)
		log.Printf("problem2: %f", n)
		// log.Printf("problem2: job=%d slot=%03d count=%05d rand=%f", j, slot, inx, r)
	}

	log.Printf("problem2: finished -------------------------------------------")
}

func printRandom2(slot int, wg *sync.WaitGroup, jobs <-chan int, nums chan<- float32) {

	for inx := 0; inx < 10; inx++ {

		select {
		case _, ok := <-jobs:
			if ok {
				r := rand.Float32()
				// log.Printf("problem2: job=%d slot=%03d count=%05d rand=%f", j, slot, inx, r)
				nums <- r
				wg.Done()
			} else {
				return
			}
		default:
			log.Printf("no value at slot=%03d, moving on", slot)
		}
	}
}
