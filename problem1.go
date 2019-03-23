package main

import (
	"log"
	"math/rand"
	"sync"
)

func problem1() {

	log.Printf("problem1: started --------------------------------------------")

	// jobs channel is used for communicating between
	// go routines to make sure 100 jobs in total will be
	// processed despite the number of loops or go routines
	jobs := make(chan int, 100)
	for j := 1; j <= 100; j++ {
		jobs <- j
	}

	// channel can be closed once the jobs are in place
	close(jobs)

	// a wait group is used instead of time.Sleep()
	var wg sync.WaitGroup

	// wait group will have 100 jobs
	wg.Add(100)

	for inx := 0; inx < 10; inx++ {
		go printRandom1(inx, &wg, jobs)
	}

	// wait group will wait for all jobs to be finished
	wg.Wait()

	log.Printf("problem1: finished --------------------------------------------")
}

func printRandom1(slot int, wg *sync.WaitGroup, jobs <-chan int) {

	//
	// Do not change 25 into 10!
	//

	for inx := 0; inx < 25; inx++ {
		select {
		case j, ok := <-jobs:
			if ok {
				log.Printf("problem1: job=%d slot=%03d count=%05d rand=%f", j, slot, inx, rand.Float32())
				wg.Done()
			} else {
				log.Printf("channel closed at slot=%03d", slot)
				return
			}
		default:
			log.Printf("no value at slot=%03d, moving on", slot)
		}
	}
}
