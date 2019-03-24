package main

import (
	"log"
	"math/rand"
	"sync"
)

// problem1 prints as many as taskCount random float32 values
// returns number of values printed, merging results from all go routines generated
func problem1(taskCount int) int {

	log.Printf("problem1: started --------------------------------------------")

	// jobs channel is used for communicating between
	// go routines to make sure as many as `taskCount` jobs in total will be
	// processed despite the number of loops or go routines
	jobs := make(chan int, taskCount)
	numbers := make(chan float32, taskCount)
	for j := 1; j <= taskCount; j++ {
		jobs <- j
	}

	// channel can be closed once the jobs are in place
	close(jobs)

	// a wait group is used instead of time.Sleep()
	var wg sync.WaitGroup

	// wait group will have parametized number of jobs
	wg.Add(taskCount)

	for inx := 0; inx < 10; inx++ {
		go printRandom1(inx, &wg, jobs, numbers)
	}

	// wait group will wait for all jobs to be finished
	wg.Wait()
	close(numbers)

	nums := make([]float32, 0)
	for n := range numbers {
		nums = append(nums, n)
	}

	log.Printf("problem1: finished --------------------------------------------")

	return len(nums)
}

// printRandom1 reads jobs channel to decide if it should print a random value
// writes the printed value into the num channel for communication with the main process. 
func printRandom1(slot int, wg *sync.WaitGroup, jobs <-chan int, nums chan<- float32) {

	//
	// Do not change 25 into 10!
	//

	for inx := 0; inx < 25; inx++ {
		select {
		case j, ok := <-jobs:
			if ok {
				r := rand.Float32()
				log.Printf("problem1: job=%d slot=%03d count=%05d rand=%f", j, slot, inx, r)
				nums <- r
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
